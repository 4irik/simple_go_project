package main

import (
	"slices"
	"strings"
)

type SliceWhiteList struct {
	names []string
}

type WhiteList interface {
	Add(string) bool
	IsExist(string) bool
}

func (wl *SliceWhiteList) Add(newName string) bool {
	newName = strings.TrimSpace(newName)
	if newName == "" || wl.IsExist(newName) {
		return false
	}
	wl.names = append(wl.names, newName)
	return true
}

func (wl *SliceWhiteList) IsExist(name string) bool {
	return slices.Contains(wl.names, name)
}
