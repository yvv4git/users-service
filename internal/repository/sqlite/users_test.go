package sqlite

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/yvv4git/users-service/internal/domain"
	"github.com/yvv4git/users-service/internal/repository"
)

func TestUsers_Create(t *testing.T) {
	// Query for add user to storage.
	insertQuery := `INSERT INTO users \(name, email, age\) VALUES \(\?, \?, \?\)`

	// Create mock and mock-db.
	db, mock, err := sqlmock.New()
	assert.Nil(t, err)
	defer db.Close()

	// Init UsersRepository with mock-db.
	repo := NewUsersRepository(db)

	tests := []struct {
		name      string
		repo      repository.UsersRepository
		user      domain.Users
		mock      func(mock sqlmock.Sqlmock)
		want      int
		wantError bool
	}{
		{
			name: "Create new user",
			repo: repo,
			user: domain.Users{
				Name:  "Vladimir Eliseev",
				Email: "yvv4test@gmail.com",
				Age:   32,
			},
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(insertQuery).
					WithArgs("Vladimir Eliseev", "yvv4test@gmail.com", 32).
					WillReturnResult(sqlmock.NewResult(1, 1)) // LastInsert, RowsAffected
			},
			wantError: false,
		},
		{
			name: "Try add empty user",
			repo: repo,
			user: domain.Users{},
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(insertQuery).
					WithArgs("", "", 0).
					WillReturnError(repository.ErrEmptyParameters)
			},
			wantError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			//t.Parallel()

			tc.mock(mock)

			user, err := tc.repo.Create(tc.user)
			if tc.wantError {
				assert.NotEmpty(t, err.Error(), "Empty parameters")
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tc.user.Name, user.Name)
				assert.Equal(t, tc.user.Email, user.Email)
				assert.Equal(t, tc.user.Age, user.Age)
			}
		})
	}

	mock.ExpectClose()
	err = db.Close()
	assert.Nil(t, err)
}

func TestUsers_Read(t *testing.T) {
	selectQuery := `SELECT id, name, email, age FROM users WHERE id = \?`

	// Create mock and mock-db.
	db, mock, err := sqlmock.New()
	assert.Nil(t, err)
	defer db.Close()

	// Init UsersRepository with mock-db.
	repo := NewUsersRepository(db)

	// Create test cases.
	tests := []struct {
		name      string
		repo      repository.UsersRepository
		user      domain.Users
		mock      func(mock sqlmock.Sqlmock)
		want      int
		wantError bool
	}{
		{
			name: "Find user by id",
			repo: repo,
			user: domain.Users{
				ID:    1,
				Name:  "Vladimir Eliseev",
				Email: "yvv4test@gmail.com",
				Age:   32,
			},
			mock: func(mock sqlmock.Sqlmock) {
				// Fill mock-db with add new rows.
				rows := sqlmock.NewRows([]string{"id", "name", "email", "age"}).
					AddRow(1, "Vladimir Eliseev", "yvv4test@gmail.com", 32)

				// Expect select query.
				mock.ExpectQuery(selectQuery).
					WithArgs(1).
					WillReturnRows(rows)
			},
			wantError: false,
		},
		{
			name: "Find user by id",
			repo: repo,
			user: domain.Users{ID: 0},
			mock: func(mock sqlmock.Sqlmock) {
				// Expect select query.
				mock.ExpectQuery(selectQuery).
					WithArgs(0).
					WillReturnError(repository.ErrUserNotFound)
			},
			wantError: true,
		},
	}

	// Use test cases.
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			//t.Parallel()

			tc.mock(mock)

			user, err := tc.repo.Read(tc.user.ID)
			if tc.wantError {
				assert.NotEmpty(t, err.Error(), "User not found")
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tc.user.Name, user.Name)
				assert.Equal(t, tc.user.Email, user.Email)
				assert.Equal(t, tc.user.Age, user.Age)
			}
		})
	}

	// Wait close connection.
	mock.ExpectClose()
	err = db.Close()
	assert.Nil(t, err)
}

func TestUsers_Update(t *testing.T) {
	updateQuery := `UPDATE users SET name = \?, email = \?, age = \? WHERE id = \?`

	// Create mock and mock-db.
	db, mock, err := sqlmock.New()
	assert.Nil(t, err)
	defer db.Close()

	// Init UsersRepository with mock-db.
	repo := NewUsersRepository(db)

	// Init test cases.
	tests := []struct {
		name      string
		repo      repository.UsersRepository
		user      domain.Users
		mock      func(mock sqlmock.Sqlmock)
		want      int
		wantError bool
	}{
		{
			name: "Update user by id",
			repo: repo,
			user: domain.Users{
				ID:    1,
				Name:  "Vladimir Eliseev",
				Email: "yvv4test@gmail.com",
				Age:   32,
			},
			mock: func(mock sqlmock.Sqlmock) {
				// Expect update query.
				mock.ExpectExec(updateQuery).
					WithArgs("Vladimir Eliseev", "yvv4test@gmail.com", 32, 1).
					WillReturnResult(sqlmock.NewResult(1, 1)) // LastInsert, RowsAffected
			},
			wantError: false,
		},
		{
			name: "Update the user when it doesn't exist",
			repo: repo,
			user: domain.Users{
				ID:    0,
				Name:  "Vladimir Eliseev",
				Email: "yvv4test@gmail.com",
				Age:   32,
			},
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(updateQuery).
					WithArgs("Vladimir Eliseev", "yvv4test@gmail.com", 32, 0).
					WillReturnError(repository.ErrUserNotFound)
			},
			wantError: true,
		},
	}

	// Use test cases.
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			//t.Parallel()

			tc.mock(mock)

			err := tc.repo.Update(tc.user)
			if tc.wantError {
				assert.NotEmpty(t, err.Error(), "User not found")
			} else {
				assert.Nil(t, err)
			}
		})
	}

	// Wait close connection.
	mock.ExpectClose()
	err = db.Close()
	assert.Nil(t, err)
}

func TestUsers_Del(t *testing.T) {
	deleteQuery := `DELETE FROM users WHERE id = \?`

	// Create mock and mock-db.
	db, mock, err := sqlmock.New()
	assert.Nil(t, err)
	defer db.Close()

	// Init UsersRepository with mock-db.
	repo := NewUsersRepository(db)

	// Init test cases.
	tests := []struct {
		name      string
		repo      repository.UsersRepository
		user      domain.Users
		mock      func(mock sqlmock.Sqlmock)
		want      int
		wantError bool
	}{
		{
			name: "Delete user by id",
			repo: repo,
			user: domain.Users{ID: 1},
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(deleteQuery).
					WithArgs(1).
					WillReturnResult(sqlmock.NewResult(1, 1)) // LastInsert, RowsAffected
			},
			wantError: false,
		},
		{
			name: "Delete user by id when it doesn't exist",
			repo: repo,
			user: domain.Users{ID: 0},
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(deleteQuery).
					WithArgs(0).
					WillReturnError(repository.ErrUserNotFound)
			},
			wantError: true,
		},
	}

	// Use test cases.
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			//t.Parallel()

			tc.mock(mock)

			err := tc.repo.Del(tc.user.ID)
			if tc.wantError {
				assert.NotEmpty(t, err.Error(), "User not found")
			} else {
				assert.Nil(t, err)
			}
		})
	}

	// Wait close connection.
	mock.ExpectClose()
	err = db.Close()
	assert.Nil(t, err)
}
