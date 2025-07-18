package usecase

import (
	"fmt"
	"microservice/domain/dto"
	"microservice/domain/interfaces"
)

type FindByIdUseCase struct {
	BalanceRepository interfaces.IBalanceRepository
}

func NewFindByIdUseCase(balanceRepository interfaces.IBalanceRepository) *FindByIdUseCase {
	return &FindByIdUseCase{
		BalanceRepository: balanceRepository,
	}
}

func (c *FindByIdUseCase) Execute(input any) (any, error) {
	fmt.Println("input", input)
	findInput := input.(*dto.FindByIdInput)
	balance, err := c.BalanceRepository.FindById(findInput.ID)
	if err != nil {
		return nil, err
	}
	return &dto.FindByIdOutput{
		ID:        balance.ID,
		AccountID: balance.AccountID,
		Balance:   balance.Balance,
	}, nil
}

// Verificar se implementa a interface
var _ interfaces.IBalanceUseCase = (*FindByIdUseCase)(nil)
