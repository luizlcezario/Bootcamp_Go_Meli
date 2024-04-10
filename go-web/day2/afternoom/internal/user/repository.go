package users

type Repository interface {
	GetAll() ([]User, error)
	FindById(id string) (User, error)
	Store(User) (User, error)
	GetQueryParametersValids() map[string]int
}
