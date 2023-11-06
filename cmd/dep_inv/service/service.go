package service

import (
	"errors"
	"sandbox/cmd/dep_inv/repository"
)

var ErrUserNotFound = errors.New("no such user")

type Service struct {
	R repository.IRepo
}

func (s *Service) SetBalance(uid, money int) {
	s.R.Set(uid, money)
}

func (s *Service) GetBalance(uid int) (int, error) {
	money, ok := s.R.Get(uid)
	if !ok {
		return 0, ErrUserNotFound
	}
	return money, nil
}
