package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yvv4git/users-service/internal/domain"
	"github.com/yvv4git/users-service/internal/repository"
)

func TestUsersService_Create(t *testing.T) {
	// A set of tests.
	testsCases := []struct {
		name         string
		usersPayload domain.Users
		mock         func(mock *repository.MockUsersRepository, usersPayload domain.Users)
		wantError    bool
	}{
		{
			name: "Create user",
			usersPayload: domain.Users{
				Name:  "Vladimir Eliseev",
				Email: "yvv4test@gmail.com",
				Age:   32,
			},
			mock: func(mock *repository.MockUsersRepository, usersPayload domain.Users) {
				mock.On("Create", usersPayload).Return(&usersPayload, nil).Once()
			},
			wantError: false,
		},
		{
			name: "Try create empty user",
			usersPayload: domain.Users{
				Name:  "",
				Email: "",
				Age:   0,
			},
			mock: func(mock *repository.MockUsersRepository, usersPayload domain.Users) {
				mock.On("Create", usersPayload).Return(nil, repository.ErrEmptyParameters).Once()
			},
			wantError: true,
		},
	}

	for _, tc := range testsCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepository := repository.MockUsersRepository{}
			tc.mock(&mockRepository, tc.usersPayload)

			// Check function Create.
			usersService := NewUsersService(&mockRepository)
			users, err := usersService.Create(tc.usersPayload)

			if tc.wantError {
				assert.NotEmpty(t, err)
				assert.Nil(t, users)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tc.usersPayload, *users)
			}
		})
	}
}

func TestUsersService_Read(t *testing.T) {
	testCases := []struct {
		name         string
		usersPayload domain.Users
		mock         func(mock *repository.MockUsersRepository, usersPayload domain.Users)
		wantError    bool
	}{
		{
			name: "Find user by id",
			usersPayload: domain.Users{
				ID:    1,
				Name:  "Vladimir Eliseev",
				Email: "yvv4test@gmail.com",
				Age:   32,
			},
			mock: func(mock *repository.MockUsersRepository, usersPayload domain.Users) {
				mock.On("Read", usersPayload.ID).Return(&usersPayload, nil).Once()
			},
			wantError: false,
		},
		{
			name: "User not found",
			usersPayload: domain.Users{
				ID: 0,
			},
			mock: func(mock *repository.MockUsersRepository, usersPayload domain.Users) {
				mock.On("Read", usersPayload.ID).Return(nil, repository.ErrEmptyParameters).Once()
			},
			wantError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepository := repository.MockUsersRepository{}
			tc.mock(&mockRepository, tc.usersPayload)

			// Check function Read.
			usersService := NewUsersService(&mockRepository)
			users, err := usersService.Read(tc.usersPayload.ID)

			if tc.wantError {
				assert.NotEmpty(t, err)
				assert.Nil(t, users)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tc.usersPayload, *users)
			}
		})
	}
}

func TestUsers_Update(t *testing.T) {
	testCases := []struct {
		name      string
		user      domain.Users
		mock      func(mock *repository.MockUsersRepository, inputUser domain.Users)
		wantError bool
	}{
		{
			name: "Update user",
			user: domain.Users{
				ID:    1,
				Name:  "Vladimir Eliseev",
				Email: "yvv4test@gmail.com",
				Age:   777,
			},
			mock: func(mock *repository.MockUsersRepository, inputUser domain.Users) {
				mock.On("Update", inputUser).Return(nil).Once()
			},
			wantError: false,
		},
		{
			name: "User not found for update",
			user: domain.Users{
				ID:    0,
				Name:  "Vladimir Eliseev",
				Email: "yvv4test@gmail.com",
				Age:   777,
			},
			mock: func(mock *repository.MockUsersRepository, inputUser domain.Users) {
				mock.On("Update", inputUser).Return(repository.ErrUserNotFound).Once()
			},
			wantError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepository := repository.MockUsersRepository{}
			tc.mock(&mockRepository, tc.user)

			// Check function Update.
			usersService := NewUsersService(&mockRepository)
			err := usersService.Update(tc.user)

			if tc.wantError {
				assert.NotEmpty(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestUsersService_Del(t *testing.T) {
	testCases := []struct {
		name      string
		user      domain.Users
		mock      func(mock *repository.MockUsersRepository, id int64)
		wantError bool
	}{
		{
			name: "Delete user by id",
			user: domain.Users{
				ID: 1,
			},
			mock: func(mock *repository.MockUsersRepository, id int64) {
				mock.On("Del", id).Return(nil).Once()
			},
			wantError: false,
		},
		{
			name: "No found user when delete",
			user: domain.Users{
				ID: 0,
			},
			mock: func(mock *repository.MockUsersRepository, id int64) {
				mock.On("Del", id).Return(nil).Once()
			},
			wantError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepository := repository.MockUsersRepository{}
			tc.mock(&mockRepository, tc.user.ID)

			// Check function Del.
			usersService := NewUsersService(&mockRepository)
			err := usersService.Del(tc.user.ID)

			if tc.wantError {
				assert.NotEmpty(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
