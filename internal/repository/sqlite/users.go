package sqlite

import (
	"database/sql"

	"github.com/yvv4git/users-service/internal/domain"
	"github.com/yvv4git/users-service/internal/repository"
)

// UsersRepository is used as repository for users entity.
type UsersRepository struct {
	db *sql.DB
}

// NewUsersRepository is used as contructor for UsersRepository.
func NewUsersRepository(db *sql.DB) *UsersRepository {
	return &UsersRepository{db: db}
}

// Create users entity in data store.
func (r *UsersRepository) Create(user domain.Users) (result *domain.Users, err error) {
	res, err := r.db.Exec("INSERT INTO users (name, email, age) VALUES (?, ?, ?)", user.Name, user.Email, user.Age)
	if err != nil {
		return nil, err
	}

	lastid, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	user.ID = lastid
	return &user, err
}

// Read is used for find user in data storage by id.
func (r *UsersRepository) Read(id int64) (result *domain.Users, err error) {
	var user domain.Users
	err = r.db.
		QueryRow("SELECT id, name, email, age FROM users WHERE id = ?", id).
		Scan(&user.ID, &user.Name, &user.Email, &user.Age)
	if err != nil {
		return nil, err
	}

	return &user, err
}

// Update is used for update user in data storage by id.
func (r *UsersRepository) Update(user domain.Users) (err error) {
	result, err := r.db.Exec(
		"UPDATE users SET name = ?, email = ?, age = ? WHERE id = ?",
		user.Name,
		user.Email,
		user.Age,
		user.ID,
	)
	if err != nil {
		return err
	}

	cnt, err := result.RowsAffected()
	if cnt <= 0 {
		return repository.ErrUserNotFound
	}

	return err
}

// Del is used for delete user from data storage by id.
func (r *UsersRepository) Del(id int64) (err error) {
	result, err := r.db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return err
	}

	cnt, err := result.RowsAffected()
	if cnt <= 0 {
		return repository.ErrUserNotFound
	}

	return err
}
