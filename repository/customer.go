package repository

type CustomerRepository interface {
	GetAll()
	GetById(int)
}
