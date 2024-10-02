package user_repository

import (
	"errors"
	"io/fs"
)

type FileBasedUserRepository struct {
	users []string
	file  fs.File
}

func NewFileBasedUserRepository(fsys fs.FS, fileName string) (*FileBasedUserRepository, error) {
	file, err := fsys.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return &FileBasedUserRepository{*new([]string), file}, nil
}

func (r *FileBasedUserRepository) Add(user string) error {
	return errors.New("todo")
}

func (r *FileBasedUserRepository) Delete(user string) error {
	return errors.New("todo")
}

func (r *FileBasedUserRepository) IsExist(user string) (bool, error) {
	return false, errors.New("todo")
}
