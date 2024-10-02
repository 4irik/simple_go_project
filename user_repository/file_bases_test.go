package user_repository

import (
	"io/fs"
	"testing"

	"github.com/stretchr/testify/suite"
)

type FileBasedUserRepositoryTestSuite struct {
	suite.Suite
}

type mockedFile struct {
	isOpen  *bool
	content string
}

func (f mockedFile) Stat() (fs.FileInfo, error) {
	return nil, nil
}
func (f mockedFile) Read([]byte) (int, error) {
	return 0, nil
}
func (f mockedFile) Close() error {
	*f.isOpen = false

	return nil
}

type mockedFS struct {
	file mockedFile
	err  error
}

func (mfs mockedFS) Open(name string) (fs.File, error) {
	if mfs.err != nil {
		return nil, mfs.err
	}

	*mfs.file.isOpen = true
	return &mfs.file, nil
}

// func (suite *FileBasedUserRepositoryTestSuite) TestFileError() {
// 	fs := mockedFS{mockedFile{}, errors.New("Какя-то проблема с файлом")}

// 	repo, err := NewFileBasedUserRepository(fs, "some_file_name.txt")
// 	suite.Nil(repo, "Репозиторий не должен быть создан")
// 	suite.Equal(errors.New("Какя-то проблема с файлом"), err)
// }

// func (suite *FileBasedUserRepositoryTestSuite) TestFilePermissionFail() {
// 	suite.FailNow("todo")
// }

// func (suite *FileBasedUserRepositoryTestSuite) TestNew() {
// 	fs := mockedFS{mockedFile{}, nil}
// 	repo, err := NewFileBasedUserRepository(fs, "some_file_name.txt")

// 	suite.Nil(err, "Ошибок при создании репозитория не должно быть")
// 	suite.Equal(fs.file, repo.file)
// }

func (suite *FileBasedUserRepositoryTestSuite) TestCloseFile() {
	isOpen := false
	fs := mockedFS{mockedFile{&isOpen, ""}, nil}
	repo, _ := NewFileBasedUserRepository(&fs, "some_file_name.txt")
	suite.True(isOpen, "Файл не был открыт")
	suite.IsType(&FileBasedUserRepository{}, repo)
	repo = new(FileBasedUserRepository)
	suite.False(isOpen, "Файл не был закрыт")
}

// func (suite *FileBasedUserRepositoryTestSuite) TestAddSuccess() {
// 	file := mockedFile{false, "Иван\nПётр\nСергей\nНиколай\n"}
// 	fs := mockedFS{file, nil}

// 	repo, err := NewFileBasedUserRepository(fs, "some_file_name.txt")
// }

// func (suite *FileBasedUserRepositoryTestSuite) TestIsExist() {
// 	tests := map[string]struct {
// 		name       string
// 		repository FileBasedUserRepository
// 		expected   bool
// 	}{
// 		"Имени нет в списке":   {"Иван", FileBasedUserRepository{users: []string{"Пётр", "Сергей"}}, false},
// 		"Список пуст":          {"Иван", FileBasedUserRepository{users: []string{}}, false},
// 		"Имя из пустой строки": {"", FileBasedUserRepository{users: []string{"Пётр", "Сергей"}}, false},
// 		"Имя есть в списке":    {"Иван", FileBasedUserRepository{users: []string{"Пётр", "Иван", "Сергей"}}, true},
// 	}

// 	for name, tc := range tests {
// 		suite.Run(name, func() {
// 			got, _ := tc.repository.IsExist(tc.name)
// 			suite.Equal(got, tc.expected, "Ответ не соответсвует ожиданию")
// 		})
// 	}
// }

// func (suite *FileBasedUserRepositoryTestSuite) TestAdd() {
// 	tests := map[string]struct {
// 		name         string
// 		repository   FileBasedUserRepository
// 		err          error
// 		expectedList []string
// 	}{
// 		"Новое имя": {"Иван", FileBasedUserRepository{users: []string{"Пётр", "Сергей"}}, nil, []string{"Пётр", "Сергей", "Иван"}},
// 		"Добавляем в пустой список": {"Иван", FileBasedUserRepository{users: []string{}}, nil, []string{"Иван"}},
// 		"Пустое имя":                {"", FileBasedUserRepository{users: []string{"Пётр", "Сергей"}}, errors.New("Нельзя добавить пустую строку"), []string{"Пётр", "Сергей"}},
// 		"Имя уже есть в списке":     {"Иван", FileBasedUserRepository{users: []string{"Пётр", "Иван", "Сергей"}}, errors.New("Имя уже есть в списке"), []string{"Пётр", "Иван", "Сергей"}},
// 		"`trim` на имя":             {" Николай ", FileBasedUserRepository{users: []string{"Пётр", "Иван", "Сергей"}}, nil, []string{"Пётр", "Иван", "Сергей", "Николай"}},
// 		"Имя из пробелов":           {"   ", FileBasedUserRepository{users: []string{"Пётр", "Сергей"}}, errors.New("В имени должен быть хоть один не пробельный символ"), []string{"Пётр", "Сергей"}},
// 	}
// 	for name, tc := range tests {
// 		suite.Run(name, func() {
// 			got := tc.repository.Add(tc.name)
// 			suite.Equal(got, tc.err, "Ответ не соответсвует ожиданию")
// 			suite.Equal(tc.expectedList, tc.repository.users, "Списко имён не соответсвует ожиданию")
// 		})
// 	}
// }

// func (suite *FileBasedUserRepositoryTestSuite) TestDelete() {
// 	tests := map[string]struct {
// 		delName      string
// 		repository   FileBasedUserRepository
// 		err          error
// 		expectedList []string
// 	}{
// 		"Пустое имя":         {"", FileBasedUserRepository{users: []string{"Иван", "Пётр"}}, nil, []string{"Иван", "Пётр"}},
// 		"Имени нет в списке": {"Антон", FileBasedUserRepository{users: []string{"Иван", "Пётр"}}, nil, []string{"Иван", "Пётр"}},
// 		"Имя удалено":        {"Иван", FileBasedUserRepository{users: []string{"Иван", "Пётр"}}, nil, []string{"Пётр"}},
// 		"`trim` на имя":      {" Пётр ", FileBasedUserRepository{users: []string{"Иван", "Пётр"}}, nil, []string{"Иван"}},
// 	}

// 	for name, tc := range tests {
// 		suite.Run(name, func() {
// 			got := tc.repository.Delete(tc.delName)
// 			suite.Equal(got, tc.err, "Ответ не соответсвует ожиданию")
// 			suite.Equal(tc.expectedList, tc.repository.users, "Списко имён не соответсвует ожиданию")
// 		})
// 	}
// }

func TestFileBasedUserRepositorySuite(t *testing.T) {
	suite.Run(t, new(FileBasedUserRepositoryTestSuite))
}
