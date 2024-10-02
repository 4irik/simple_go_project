package user_repository

import (
	"fmt"
	"slices"
)

type InMemoryUserRepository struct {
	users []string
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return new(InMemoryUserRepository)
}

func (r *InMemoryUserRepository) Add(user string) error {
	if exist, _ := r.IsExist(user); exist {
		return fmt.Errorf("пользователь с именем \"%s\" уже есть в списке", user)
	}

	r.users = append(r.users, user)

	return nil
}

func (r *InMemoryUserRepository) Delete(user string) error {
	i := slices.IndexFunc(r.users, func(s string) bool {
		return s == user
	})

	if i >= 0 {
		r.users = append(r.users[0:i], r.users[(i+1):]...)
	}

	return nil
}

func (r *InMemoryUserRepository) IsExist(user string) (bool, error) {
	return slices.Contains(r.users, user), nil
}
