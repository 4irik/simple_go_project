package main

import (
	"fmt"
	"reflect"
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

func TestDelete(t *testing.T) {
	type TestPair struct {
		delName        string
		before         SliceWhiteList
		expectedResult bool
		after          SliceWhiteList
	}

	testpair := []TestPair{
		{"", SliceWhiteList{names: []string{"Иван", "Пётр"}}, false, SliceWhiteList{names: []string{"Иван", "Пётр"}}},
		{"Антон", SliceWhiteList{names: []string{"Иван", "Пётр"}}, false, SliceWhiteList{names: []string{"Иван", "Пётр"}}},
		{"Иван", SliceWhiteList{names: []string{"Иван", "Пётр"}}, true, SliceWhiteList{names: []string{"Пётр"}}},
		{" Пётр ", SliceWhiteList{names: []string{"Иван", "Пётр"}}, true, SliceWhiteList{names: []string{"Иван"}}},
	}

	for _, pair := range testpair {
		v := pair.before.Delete(pair.delName)
		if v != pair.expectedResult || !reflect.DeepEqual(pair.before, pair.after) {
			t.Error(
				"For name", fmt.Sprintf("\"%s\"", pair.delName),
				"before ", pair.before,
				"after ", pair.after,
				"expected result", pair.expectedResult,
				"got result", v,
			)
		}
	}
}
