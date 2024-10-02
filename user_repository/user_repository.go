package user_repository

type UserRepositoryInterface interface {
	Add(string) error
	Delete(string) error
	IsExist(string) (bool, error)
}
