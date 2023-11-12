package service

import (
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sherpalden/sal-dist/entity"
	"github.com/sherpalden/sal-dist/repository"
)

type SalaryService interface {
	Distribute(*ethclient.Client) error
}

type salaryService struct {
	Sender       *entity.Sender
	EmployeeRepo repository.EmployeeRepository
}

func NewSalaryService(sender *entity.Sender, empRepo repository.EmployeeRepository) SalaryService {
	return salaryService{
		Sender:       sender,
		EmployeeRepo: empRepo,
	}
}

func (srv salaryService) Distribute(cl *ethclient.Client) error {
	employees, err := srv.EmployeeRepo.List()
	if err != nil {
		return err
	}
	for _, emp := range employees {
		senderBalance, err := srv.Sender.GetBalance(cl)
		if err != nil {
			return err
		}

		bigIntSenderBalance := senderBalance.BigInt()
		if bigIntSenderBalance.Cmp(emp.Salary) == -1 {
			return errors.New("insufficient sender balance")
		}

		tx := entity.Transaction{
			Sender: srv.Sender,
			ToAccount: entity.Account{
				Address: common.HexToAddress(emp.WalletAddress),
			},
			Value: emp.Salary,
		}
		if err := tx.Execute(cl); err != nil {
			return err
		}
		fmt.Printf("\nSent amount %v to employee with id: %v and name: %v\n", emp.Salary, emp.ID, emp.Name)
	}
	return nil
}
