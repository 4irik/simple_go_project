package main

import (
	"fmt"
	"testing"
)

type testpair struct {
	name      string
	whiteList SliceWhiteList
	expected  bool
}

var testsIsExist = []testpair{
	{"Иван", SliceWhiteList{names: []string{"Пётр", "Сергей"}}, false},
	{"Иван", SliceWhiteList{names: []string{}}, false},
	{"", SliceWhiteList{names: []string{"Пётр", "Сергей"}}, false},
	{"Иван", SliceWhiteList{names: []string{"Пётр", "Иван", "Сергей"}}, true},
}

var testAdd = []testpair{
	{"Иван", SliceWhiteList{names: []string{"Пётр", "Сергей"}}, true},
	{"Иван", SliceWhiteList{names: []string{}}, true},
	{"", SliceWhiteList{names: []string{"Пётр", "Сергей"}}, false},
	{"Иван", SliceWhiteList{names: []string{"Пётр", "Иван", "Сергей"}}, false},
}

func TestIsExist(t *testing.T) {
	for _, pair := range testsIsExist {
		v := pair.whiteList.IsExist(pair.name)
		if v != pair.expected {
			t.Error(
				"For name", pair.name,
				"whiteList ", pair.whiteList,
				"expected", pair.expected,
				"got", v,
			)
		}
	}
}

func TestAdd(t *testing.T) {
	for _, pair := range testAdd {
		v := pair.whiteList.Add(pair.name)
		if v != pair.expected {
			t.Error(
				"For name", fmt.Sprintf("\"%s\"", pair.name),
				"whiteList ", pair.whiteList,
				"expected", pair.expected,
				"got", v,
			)
		}
	}
}
