package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type WhiteListTestSuite struct {
	suite.Suite
}

func (suite *WhiteListTestSuite) TestIsExist() {
	tests := map[string]struct {
		name      string
		whiteList SliceWhiteList
		expected  bool
	}{
		"Имени нет в списке":   {"Иван", SliceWhiteList{names: []string{"Пётр", "Сергей"}}, false},
		"Список пуст":          {"Иван", SliceWhiteList{names: []string{}}, false},
		"Имя из пустой строки": {"", SliceWhiteList{names: []string{"Пётр", "Сергей"}}, false},
		"Имя есть в списке":    {"Иван", SliceWhiteList{names: []string{"Пётр", "Иван", "Сергей"}}, true},
	}

	for name, tc := range tests {
		suite.Run(name, func() {
			got := tc.whiteList.IsExist(tc.name)
			suite.Equal(got, tc.expected, "Ответ не соответсвует ожиданию")
		})
	}
}

func (suite *WhiteListTestSuite) TestAdd() {
	tests := map[string]struct {
		name             string
		whiteList        SliceWhiteList
		expectedResponse bool
		expectedList     []string
	}{
		"Новое имя": {"Иван", SliceWhiteList{names: []string{"Пётр", "Сергей"}}, true, []string{"Пётр", "Сергей", "Иван"}},
		"Добавляем в пустой список": {"Иван", SliceWhiteList{names: []string{}}, true, []string{"Иван"}},
		"Пустое имя":                {"", SliceWhiteList{names: []string{"Пётр", "Сергей"}}, false, []string{"Пётр", "Сергей"}},
		"Имя уже есть в списке":     {"Иван", SliceWhiteList{names: []string{"Пётр", "Иван", "Сергей"}}, false, []string{"Пётр", "Иван", "Сергей"}},
		"`trim` на имя":             {" Николай ", SliceWhiteList{names: []string{"Пётр", "Иван", "Сергей"}}, true, []string{"Пётр", "Иван", "Сергей", "Николай"}},
	}
	for name, tc := range tests {
		suite.Run(name, func() {
			got := tc.whiteList.Add(tc.name)
			suite.Equal(got, tc.expectedResponse, "Ответ не соответсвует ожиданию")
			suite.Equal(tc.expectedList, tc.whiteList.names, "Списко имён не соответсвует ожиданию")
		})
	}
}

func (suite *WhiteListTestSuite) TestDelete() {
	tests := map[string]struct {
		delName        string
		whiteList      SliceWhiteList
		expectedResult bool
		expectedList   []string
	}{
		"Пустое имя":         {"", SliceWhiteList{names: []string{"Иван", "Пётр"}}, false, []string{"Иван", "Пётр"}},
		"Имени нет в списке": {"Антон", SliceWhiteList{names: []string{"Иван", "Пётр"}}, false, []string{"Иван", "Пётр"}},
		"Имя удалено":        {"Иван", SliceWhiteList{names: []string{"Иван", "Пётр"}}, true, []string{"Пётр"}},
		"`trim` на имя":      {" Пётр ", SliceWhiteList{names: []string{"Иван", "Пётр"}}, true, []string{"Иван"}},
	}

	for name, tc := range tests {
		suite.Run(name, func() {
			got := tc.whiteList.Delete(tc.delName)
			suite.Equal(got, tc.expectedResult, "Ответ не соответсвует ожиданию")
			suite.Equal(tc.expectedList, tc.whiteList.names, "Списко имён не соответсвует ожиданию")
		})
	}
}

func (suite *WhiteListTestSuite) TestNames() {
	tests := map[string]struct {
		whiteList      SliceWhiteList
		expectedResult map[int]string
	}{
		"Пустой список":    {SliceWhiteList{names: []string{}}, map[int]string{}},
		"Не пустой список": {SliceWhiteList{names: []string{"Иван", "Пётр"}}, map[int]string{0: "Иван", 1: "Пётр"}},
	}

	for name, tc := range tests {
		suite.Run(name, func() {
			got := make(map[int]string)
			for k, name := range tc.whiteList.Names() {
				got[k] = name
			}
			suite.Equal(tc.expectedResult, got, "Списко имён не соответсвует ожиданию")
		})
	}
}

func TestWhitelistSuite(t *testing.T) {
	suite.Run(t, new(WhiteListTestSuite))
}
