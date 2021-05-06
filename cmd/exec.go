package cmd

import (
	"context"
	"github.com/aka-achu/gRPC-gw/controller"
	"github.com/aka-achu/gRPC-gw/logging"
	"github.com/aka-achu/gRPC-gw/proto/user"
	"github.com/aka-achu/gRPC-gw/repo"
	"github.com/google/uuid"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"net"
	"net/http"
)

const (
	RPCServerAddress = "0.0.0.0:8080"
	GatewayAddress   = "0.0.0.0:8081"
)

func Run() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", RPCServerAddress)
	if err != nil {
		logging.Error.Fatalf("Failed to listen to the address %s. Err-%s \n", RPCServerAddress, err.Error())
	}

	svr := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_auth.UnaryServerInterceptor(func(ctx context.Context) (context.Context, error) {
				return context.WithValue(ctx, "traceID", uuid.New().String()), nil
			}),
		)),
	)
	db := repo.NewDB()
	db.PopulateRecords()
	user.RegisterFetchServer(
		svr,
		controller.NewUserController(repo.NewUserRepo(db)),
	)

	logging.Info.Printf("Serving gRPC on %s", RPCServerAddress)
	go func() {
		logging.Error.Fatalf("Failed to start the RPC server. Err-%s \n", svr.Serve(lis).Error())
	}()

	conn, err := grpc.DialContext(
		context.Background(),
		RPCServerAddress,
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		logging.Error.Fatalf("Failed to dial up the RPC server. Err-%s \n", err.Error())
	}

	gwmux := runtime.NewServeMux()
	err = user.RegisterFetchHandler(context.Background(), gwmux, conn)
	if err != nil {
		logging.Error.Fatalf("Failed to register the gateway. Err-%s \n", err.Error())
	}

	gwServer := &http.Server{
		Addr:    GatewayAddress,
		Handler: gwmux,
	}
	logging.Info.Printf("Serving gateway on %s", GatewayAddress)
	logging.Error.Fatalf("Failed to start the gateway service. Err-%s \n",gwServer.ListenAndServe())
}
