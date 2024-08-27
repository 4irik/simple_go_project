package main

import (
	"iter"
	"slices"
	"strings"
)

type SliceWhiteList struct {
	names []string
}

type WhiteList interface {
	Add(string) bool
	IsExist(string) bool
	Names() iter.Seq2[int, string]
}

func NewSliceWhiteList() *SliceWhiteList {
	return new(SliceWhiteList)
}

func (wl *SliceWhiteList) Add(newName string) bool {
	newName = strings.TrimSpace(newName)
	if newName == "" || wl.IsExist(newName) {
		return false
	}
	wl.names = append(wl.names, newName)
	return true
}

func (wl *SliceWhiteList) Delete(name string) bool {
	name = strings.TrimSpace(name)
	if name == "" || !wl.IsExist(name) {
		return false
	}
	i := slices.IndexFunc(wl.names, func(s string) bool {
		return s == name
	})

	head := wl.names[0:i]
	tail := wl.names[(i + 1):]
	wl.names = append(head, tail...)

	return true
}

func (wl *SliceWhiteList) IsExist(name string) bool {
	return slices.Contains(wl.names, name)
}

func (wl *SliceWhiteList) Names() iter.Seq2[int, string] {
	return func(yield func(int, string) bool) {
		for key, name := range wl.names {
			if !yield(key, name) {
				return
			}
		}
	}
}
