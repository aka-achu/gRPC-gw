package repo

import (
	"context"
	"github.com/aka-achu/gRPC-gw/model"
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
		return u, nil
	} else {
		return nil, model.ErrUserDoesNotExist
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
			users = append(users, u)
		}
	}
	return users, nil
}

func (r *userRepo) AddUser(
	ctx context.Context,
	u *user.User,
) (
	int32,
	error,
) {
	u.Id = int32(r.db.Size() + 1)

	r.db.l.Lock()
	defer r.db.l.Unlock()

	r.db.store[u.Id] = u
	return u.Id, nil
}
