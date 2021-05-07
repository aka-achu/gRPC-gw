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
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
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

	// Create a gRPC server object with middleware
	svr := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_auth.UnaryServerInterceptor(func(ctx context.Context) (context.Context, error) {
				traceID := uuid.New().String()
				if meta, ok := metadata.FromIncomingContext(ctx); ok {
					logging.Info.Printf("RequestTraceID-%s ClientIP-%v UserAgent-%v GatewayUserAgent-%v",
						traceID,
						meta["x-forwarded-for"],
						meta["user-agent"],
						meta["grpcgateway-user-agent"],
					)
				}
				return context.WithValue(ctx, "traceID", traceID), nil
			}),
		)),
	)
	reflection.Register(svr)
	// Create a new database object
	db := repo.NewDB()
	// Populate some sample user details
	db.PopulateRecords()

	// Attach the Fetch service to the server
	user.RegisterFetchServer(
		svr,
		controller.NewUserController(repo.NewUserRepo(db)),
	)

	// Serve gRPC Server
	logging.Info.Printf("Serving gRPC on %s", RPCServerAddress)
	go func() {
		logging.Error.Fatalf("Failed to start the RPC server. Err-%s \n", svr.Serve(lis).Error())
	}()

	// Create a client connection to the gRPC server
	// The gRPC-gateway will proxy the requests to this server
	conn, err := grpc.DialContext(
		context.Background(),
		RPCServerAddress,
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		logging.Error.Fatalf("Failed to dial up the RPC server. Err-%s \n", err.Error())
	}

	// Create a gateway mux
	gwmux := runtime.NewServeMux()
	err = user.RegisterFetchHandler(context.Background(), gwmux, conn)
	if err != nil {
		logging.Error.Fatalf("Failed to register the gateway. Err-%s \n", err.Error())
	}

	gwServer := &http.Server{
		Addr:    GatewayAddress,
		Handler: gwmux,
	}

	// Serve gateway
	logging.Info.Printf("Serving gateway on %s", GatewayAddress)
	logging.Error.Fatalf("Failed to start the gateway service. Err-%s \n", gwServer.ListenAndServe())
}
