package repository

import (
	"github.com/sherpalden/sal-dist/entity"
	"github.com/sherpalden/sal-dist/store"
)

type EmployeeRepository interface {
	Add(entity.Employee) error
	List() ([]entity.Employee, error)
}

type employeeRepository struct {
	Store store.Store
}

func NewEmployeeRepository(store store.Store) EmployeeRepository {
	return employeeRepository{Store: store}
}

func (repo employeeRepository) Add(emp entity.Employee) error {
	return repo.Store.Create(emp)
}

func (repo employeeRepository) List() ([]entity.Employee, error) {
	var employees []entity.Employee
	return employees, repo.Store.FindAll(&employees)
}
