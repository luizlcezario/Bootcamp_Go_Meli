package users

type Repository interface {
	GetAll() ([]User, error)
	FindById(id string) (User, error)
	FindByEmail(email string) (User, error)
	Store(us User) (User, error)
	Update(dst User, src User) User
	Delete(us User) error
	GetQueryParametersValids() map[string]int
}
