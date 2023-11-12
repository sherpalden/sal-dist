package service

import (
	"errors"

	"github.com/sherpalden/sal-dist/entity"
	"github.com/sherpalden/sal-dist/repository"
)

type EmployeeService interface {
	Create(entity.Employee) error
	GetAll() ([]entity.Employee, error)
}

type employeeService struct {
	Repo repository.EmployeeRepository
}

func NewEmployeeService(repo repository.EmployeeRepository) EmployeeService {
	return employeeService{
		Repo: repo,
	}
}

func (srv employeeService) Create(emp entity.Employee) error {
	employees, err := srv.Repo.List()
	if err != nil {
		return err
	}
	for _, empItem := range employees {
		if emp.ID == empItem.ID || emp.WalletAddress == empItem.WalletAddress {
			return errors.New("employee already exits with given id and wallet-address")
		}
	}
	return srv.Repo.Add(emp)
}

func (srv employeeService) GetAll() ([]entity.Employee, error) {
	return srv.Repo.List()
}
