package repository

import "github.com/yvv4git/users-service/internal/domain"

//go:generate mockery --name=UsersRepository --output=. --filename=mock.go --inpackage
// UsersRepository is a interface for working with db.
type UsersRepository interface {
	Create(user domain.Users) (result *domain.Users, err error)
	Read(id int64) (result *domain.Users, err error)
	Update(user domain.Users) (err error)
	Del(id int64) (err error)
}
