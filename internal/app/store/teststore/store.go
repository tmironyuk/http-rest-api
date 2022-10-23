package teststore

import (
	"github.com/tmironyuk/http-rest-api/internal/app/model"
	"github.com/tmironyuk/http-rest-api/internal/app/store"
)

type Store struct {
	userRepositoty *UserRepository
}

func New() *Store {
	return &Store{}
}

func (s *Store) User() store.UserRepository {
	if s.userRepositoty != nil {
		return s.userRepositoty
	}

	s.userRepositoty = &UserRepository{
		store: s,
		users: make(map[int]*model.User),
	}

	return s.userRepositoty
}
