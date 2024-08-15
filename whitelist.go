package main

import (
	"slices"
	"strings"
)

type WhiteList struct {
	names []string
}

type WhiteListInterface interface {
	Add(string) bool
	IsExist(string) bool
}

func (wl *WhiteList) Add(newName string) bool {
	newName = strings.TrimSpace(newName)
	if newName == "" || wl.IsExist(newName) {
		return false
	}
	wl.names = append(wl.names, newName)
	return true
}

func (wl *WhiteList) IsExist(name string) bool {
	return slices.Contains(wl.names, name)
}
