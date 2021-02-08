package sqlite

import (
	"database/sql"

	"github.com/yvv4git/users-service/domain"
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

// Update ...
func (r *UsersRepository) Update(user domain.Users) (err error) {
	return err
}

// Del ...
func (r *UsersRepository) Del(id int64) (err error) {
	return err
}
