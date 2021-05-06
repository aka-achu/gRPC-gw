package controller

import (
	"context"
	"github.com/aka-achu/gRPC-gw/logging"
	"github.com/aka-achu/gRPC-gw/model"
	userpb "github.com/aka-achu/gRPC-gw/proto/user"
)

type userController struct {
	userpb.UnimplementedFetchServer
	repo model.UserRepo
}

func NewUserController(r model.UserRepo) *userController {
	return &userController{repo: r}
}

func (c *userController) FetchUserByID(
	ctx context.Context,
	req *userpb.FetchUserByIDRequest,
) (
	*userpb.FetchUserByIDResponse,
	error,
) {
	requestID := ctx.Value("traceID").(string)
	if user, err := c.repo.FetchUserByID(ctx, req.GetId()); err != nil {
		logging.Error.Printf("TraceID-%s Failed to fetch user detail by ID. Err-%s \n", requestID, err.Error())
		return nil, err
	} else {
		return &userpb.FetchUserByIDResponse{User: user}, nil
	}
}

func (c *userController) FetchUsers(
	ctx context.Context,
	req *userpb.FetchUsersRequest,
) (
	*userpb.FetchUsersResponse,
	error,
) {
	requestID := ctx.Value("traceID").(string)
	if users, err := c.repo.FetchUsers(ctx, req.GetId()); err != nil {
		logging.Error.Printf("TraceID-%s Failed to fetch user details by IDs. Err-%s \n", requestID, err.Error())
		return nil, err
	} else {
		return &userpb.FetchUsersResponse{Users: users}, nil
	}
}
