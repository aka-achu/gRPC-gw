package controller

import (
	"context"
	"github.com/aka-achu/gRPC-gw/proto/user"
	"github.com/aka-achu/gRPC-gw/repo"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestFetchUserByID(t *testing.T) {

	mockDB := repo.NewDB()
	mockUserRepo := repo.NewUserRepo(mockDB)

	mockUser := &user.User{
		Fname:   "Joey",
		City:    "LA",
		Phone:   8888888885,
		Height:  5.10,
		Married: false,
	}
	mockUserId, _ := mockUserRepo.AddUser(context.Background(), mockUser)
	mockUser.Id = mockUserId

	mockUserController := NewUserController(mockUserRepo)

	t.Run("trying to fetch an existing user detail", func(t *testing.T) {
		expectedResultSet := mockUser
		response, err := mockUserController.FetchUserByID(
			context.WithValue(context.Background(), "traceID", uuid.New().String()),
			&user.FetchUserByIDRequest{Id: mockUserId},
		)
		assert.Nil(t, err, "Failed to fetch the user by id")
		if !reflect.DeepEqual(expectedResultSet, response.User) {
			t.Errorf("Expected %v but got %v", mockUser, response.User)
		}
	})

	t.Run("trying to fetch a non-existent user detail", func(t *testing.T) {
		_, err := mockUserController.FetchUserByID(
			context.WithValue(context.Background(), "traceID", uuid.New().String()),
			&user.FetchUserByIDRequest{Id: 99},
		)
		if err == nil {
			t.Errorf("Expected user does not exist error but got %v", err)
		}
	})
}

func TestFetchUsers(t *testing.T) {
	mockDB := repo.NewDB()
	mockUserRepo := repo.NewUserRepo(mockDB)

	mockUser1 := &user.User{
		Fname:   "Ross",
		City:    "LA",
		Phone:   8888888885,
		Height:  5.10,
		Married: false,
	}
	mockUserId1, _ := mockUserRepo.AddUser(context.Background(), mockUser1)
	mockUser1.Id = mockUserId1

	mockUser2 := &user.User{
		Fname:   "Chandler",
		City:    "LA",
		Phone:   8888888885,
		Height:  5.10,
		Married: false,
	}
	mockUserId2, _ := mockUserRepo.AddUser(context.Background(), mockUser2)
	mockUser2.Id = mockUserId2

	mockUserController := NewUserController(mockUserRepo)

	t.Run("trying to fetch an existing user details", func(t *testing.T) {
		expectedResultSet := []*user.User{mockUser1, mockUser2}
		response, err := mockUserController.FetchUsers(
			context.WithValue(context.Background(), "traceID", uuid.New().String()),
			&user.FetchUsersRequest{Id: []int32{mockUserId1, mockUserId2}},
		)
		assert.Nil(t, err, "Failed to fetch the users by ids")
		if !reflect.DeepEqual(expectedResultSet, response.Users) {
			t.Errorf("Expected %v but got %v", expectedResultSet, response.Users)
		}
	})

	t.Run("trying to fetch non-existent user details", func(t *testing.T) {
		var expectedResultSet []*user.User
		response, err := mockUserController.FetchUsers(
			context.WithValue(context.Background(), "traceID", uuid.New().String()),
			&user.FetchUsersRequest{Id: []int32{99, 98}},
		)
		assert.Nil(t, err, "Failed to fetch the users by ids")
		if !reflect.DeepEqual(expectedResultSet, response.Users) {
			t.Errorf("Expected %v but got %v", expectedResultSet, response.Users)
		}
	})

	t.Run("trying to fetch no user details, with empty request", func(t *testing.T) {
		var expectedResultSet []*user.User
		response, err := mockUserController.FetchUsers(
			context.WithValue(context.Background(), "traceID", uuid.New().String()),
			&user.FetchUsersRequest{Id: []int32{}},
		)
		assert.Nil(t, err, "Failed to fetch the users by ids")
		if !reflect.DeepEqual(expectedResultSet, response.Users) {
			t.Errorf("Expected %v but got %v", expectedResultSet, response.Users)
		}
	})
}
