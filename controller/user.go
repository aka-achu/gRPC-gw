package controller

import (
	"context"
	"github.com/aka-achu/gRPC-gw/logging"
	"github.com/aka-achu/gRPC-gw/model"
	userpb "github.com/aka-achu/gRPC-gw/proto/user"
)

// userController implements user.FetchServer
type userController struct {
	userpb.UnimplementedFetchServer
	repo model.UserRepo
}

// NewUserController will initialize an instance of userController
func NewUserController(r model.UserRepo) *userController {
	return &userController{repo: r}
}

// FetchUserByID is an rpc/gateway handler to fetch user detail of the requested user id
func (c *userController) FetchUserByID(
	ctx context.Context,
	req *userpb.FetchUserByIDRequest,
) (
	*userpb.FetchUserByIDResponse,
	error,
) {
	requestID := ctx.Value("traceID").(string)
	if req.GetId() == 0 {
		logging.Error.Printf("TraceID-%s Invalid/Nil id of the requested user.\n", requestID)
		return nil, model.ErrNilUserID
	}
	if user, err := c.repo.FetchUserByID(ctx, req.GetId()); err != nil {
		logging.Error.Printf("TraceID-%s Failed to fetch user detail by ID. Err-%s \n", requestID, err.Error())
		return nil, err
	} else {
		return &userpb.FetchUserByIDResponse{User: user}, nil
	}
}

// FetchUsers is an rpc/gateway handler to fetch user details of the requested user ids
func (c *userController) FetchUsers(
	ctx context.Context,
	req *userpb.FetchUsersRequest,
) (
	*userpb.FetchUsersResponse,
	error,
) {
	requestID := ctx.Value("traceID").(string)
	if len(req.GetId()) == 0 {
		logging.Error.Printf("TraceID-%s Invalid/Nil ids of the requested user.\n", requestID)
		return nil, model.ErrNilUserID
	}
	if users, err := c.repo.FetchUsers(ctx, req.GetId()); err != nil {
		logging.Error.Printf("TraceID-%s Failed to fetch user details by IDs. Err-%s \n", requestID, err.Error())
		return nil, err
	} else {
		return &userpb.FetchUsersResponse{Users: users}, nil
	}
}
