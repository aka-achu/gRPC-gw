package model

import (
	"context"
	"github.com/aka-achu/gRPC-gw/proto/user"
)

type UserRepo interface {
	AddUser(context.Context, *user.User) (int32, error)
	FetchUserByID(context.Context, int32) (*user.User, error)
	FetchUsers(context.Context, []int32) ([]*user.User, error)
}
