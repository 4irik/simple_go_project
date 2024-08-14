package main

import "testing"

type testpair struct {
	name      string
	whiteList []string
	expected  bool
}

var tests = []testpair{
	{"Иван", []string{"Пётр", "Сергей"}, false},
	{"Иван", []string{}, false},
	{"", []string{"Пётр", "Сергей"}, false},
	{"Иван", []string{"Пётр", "Иван", "Сергей"}, true},
}

func TestInWhiteList(t *testing.T) {
	for _, pair := range tests {
		v := InWhiteList(pair.name, pair.whiteList)
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
