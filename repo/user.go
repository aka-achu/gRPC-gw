package repo

import (
	"context"
	"errors"
	"github.com/aka-achu/gRPC-gw/proto/user"
)

type userRepo struct {
	db *database
}

func NewUserRepo(db *database) *userRepo {
	return &userRepo{db: db}
}

func (r *userRepo) FetchUserByID(
	ctx context.Context,
	id int32,
) (
	*user.User,
	error,
) {
	r.db.l.Lock()
	defer r.db.l.Unlock()
	if u, ok := r.db.store[id]; ok {
		return &u, nil
	} else {
		return nil, errors.New("user with requested id does not exist")
	}
}

func (r *userRepo) FetchUsers(
	ctx context.Context,
	ids []int32,
) (
	[]*user.User,
	error,
) {
	r.db.l.Lock()
	defer r.db.l.Unlock()

	var users []*user.User
	for _, id := range ids {
		if u, ok := r.db.store[id]; ok {
			users = append(users, &u)
		}
	}
	return users, nil
}
