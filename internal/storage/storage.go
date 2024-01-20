package storage

import "github.com/romanchechyotkin/effective-mobile-test-task/internal/storage/repo"

type Collection struct {
	usersRepo *repo.Users
}

func NewCollection(users *repo.Users) *Collection {
	return &Collection{
		usersRepo: users,
	}
}

func (c Collection) Users() *repo.Users {
	return c.usersRepo
}
