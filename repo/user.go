package repo

import (
	"context"
	"github.com/aka-achu/gRPC-gw/model"
	"github.com/aka-achu/gRPC-gw/proto/user"
)

// userRepo implements model.UserRepo interface
type userRepo struct {
	db *database
}

// NewUserRepo initializes an instance of userRepo
func NewUserRepo(db *database) *userRepo {
	return &userRepo{db: db}
}

// FetchUserByID returns detail of the user having the requested id as user.id
// if the user does not exist in the database, then it will return an error
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

// FetchUsers returns details of users by iterating the requested user ids.
// If one requested id does not exist in the database, it will simple ignore that id
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

// AddUser will create an entry for the requested user
// UserID will depend on the size of the database, for autoincrement
// As there are no delete user details from the database, the user id will not clash.
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
