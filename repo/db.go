package repo

import (
	"github.com/aka-achu/gRPC-gw/proto/user"
	"sync"
)

type database struct {
	store map[int32]*user.User
	// ensuring atomic access to the store
	l     sync.Mutex
}

func NewDB() *database {
	return &database{store: make(map[int32]*user.User)}
}

func (db *database) PopulateRecords() {
	db.l.Lock()
	defer db.l.Unlock()
	db.store[1] = &user.User{
		Id:      1,
		Fname:   "John",
		City:    "New York",
		Phone:   9999999991,
		Height:  5.8,
		Married: true,
	}
	db.store[2] = &user.User{
		Id:      2,
		Fname:   "Tom",
		City:    "Paris",
		Phone:   9999999992,
		Height:  5.8,
		Married: false,
	}
	db.store[3] = &user.User{
		Id:      3,
		Fname:   "James",
		City:    "Moscow",
		Phone:   9999999993,
		Height:  5.8,
		Married: true,
	}
	db.store[4] = &user.User{
		Id:      4,
		Fname:   "Bob",
		City:    "Tokyo",
		Phone:   9999999994,
		Height:  5.8,
		Married: false,
	}
	db.store[5] = &user.User{
		Id:      5,
		Fname:   "Jose",
		City:    "Dubai",
		Phone:   9999999995,
		Height:  5.8,
		Married: false,
	}

}

func (db *database) Size() int {
	db.l.Lock()
	defer db.l.Unlock()
	return len(db.store)
}
