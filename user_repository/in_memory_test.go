package user_repository

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"
)

type InMemoryUserRepositoryTestSuite struct {
	suite.Suite
}

func (suite *InMemoryUserRepositoryTestSuite) TestNew() {
	repo := NewInMemoryUserRepository()

	suite.Len(repo.users, 0, "При создании репозитория список пользователей должен быть пуст")
}

func (suite *InMemoryUserRepositoryTestSuite) TestIsExist() {
	tests := map[string]struct {
		name       string
		repository InMemoryUserRepository
		expected   bool
	}{
		"Имени нет в списке":   {"Иван", InMemoryUserRepository{users: []string{"Пётр", "Сергей"}}, false},
		"Список пуст":          {"Иван", InMemoryUserRepository{users: []string{}}, false},
		"Имя из пустой строки": {"", InMemoryUserRepository{users: []string{"Пётр", "Сергей"}}, false},
		"Имя есть в списке":    {"Иван", InMemoryUserRepository{users: []string{"Пётр", "Иван", "Сергей"}}, true},
	}

	for name, tc := range tests {
		suite.Run(name, func() {
			got, _ := tc.repository.IsExist(tc.name)
			suite.Equal(got, tc.expected, "Ответ не соответсвует ожиданию")
		})
	}
}

func (suite *InMemoryUserRepositoryTestSuite) TestAdd() {
	tests := map[string]struct {
		name         string
		repository   InMemoryUserRepository
		err          error
		expectedList []string
	}{
		"Новое имя": {"Иван", InMemoryUserRepository{users: []string{"Пётр", "Сергей"}}, nil, []string{"Пётр", "Сергей", "Иван"}},
		"Добавляем в пустой список": {"Иван", InMemoryUserRepository{users: []string{}}, nil, []string{"Иван"}},
		"Пустое имя":                {"", InMemoryUserRepository{users: []string{"Пётр", "Сергей"}}, nil, []string{"Пётр", "Сергей", ""}},
		"Имя уже есть в списке":     {"Иван", InMemoryUserRepository{users: []string{"Пётр", "Иван", "Сергей"}}, errors.New("пользователь с именем \"Иван\" уже есть в списке"), []string{"Пётр", "Иван", "Сергей"}},
		"`trim` на имя":             {" Николай ", InMemoryUserRepository{users: []string{"Пётр", "Иван", "Сергей"}}, nil, []string{"Пётр", "Иван", "Сергей", " Николай "}},
		"Имя из пробелов":           {"   ", InMemoryUserRepository{users: []string{"Пётр", "Сергей"}}, nil, []string{"Пётр", "Сергей", "   "}},
	}
	for name, tc := range tests {
		suite.Run(name, func() {
			got := tc.repository.Add(tc.name)
			suite.Equal(got, tc.err, "Ответ не соответсвует ожиданию")
			suite.Equal(tc.expectedList, tc.repository.users, "Списко имён не соответсвует ожиданию")
		})
	}
}

func (suite *InMemoryUserRepositoryTestSuite) TestDelete() {
	tests := map[string]struct {
		delName      string
		repository   InMemoryUserRepository
		err          error
		expectedList []string
	}{
		"Пустое имя которого нет в списке": {"", InMemoryUserRepository{users: []string{"Иван", "Пётр"}}, nil, []string{"Иван", "Пётр"}},
		"Пустое имя которое есть в списке": {"", InMemoryUserRepository{users: []string{"Иван", "Пётр", ""}}, nil, []string{"Иван", "Пётр"}},
		"Имени нет в списке":               {"Антон", InMemoryUserRepository{users: []string{"Иван", "Пётр"}}, nil, []string{"Иван", "Пётр"}},
		"Имя удалено":                      {"Иван", InMemoryUserRepository{users: []string{"Иван", "Пётр"}}, nil, []string{"Пётр"}},
		"`trim` на имя":                    {" Пётр ", InMemoryUserRepository{users: []string{"Иван", "Пётр"}}, nil, []string{"Иван", "Пётр"}},
	}

	for name, tc := range tests {
		suite.Run(name, func() {
			got := tc.repository.Delete(tc.delName)
			suite.Equal(got, tc.err, "Ответ не соответсвует ожиданию")
			suite.Equal(tc.expectedList, tc.repository.users, "Списко имён не соответсвует ожиданию")
		})
	}
}

func TestInMemoryUserRepositorySuite(t *testing.T) {
	suite.Run(t, new(InMemoryUserRepositoryTestSuite))
}
