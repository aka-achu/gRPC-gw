package model

import (
	"context"
	"github.com/aka-achu/gRPC-gw/proto/user"
)

type UserRepo interface {
	FetchUserByID(context.Context, int32) (*user.User, error)
	FetchUsers(context.Context, []int32) ([]*user.User, error)
}
