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

func (c *FindByIdUseCase) Execute(input *dto.FindByIdInput) (*dto.FindByIdOutput, error) {
	fmt.Println("input", input)

	balance, err := c.BalanceRepository.FindById(input.ID)
	if err != nil {
		return nil, err
	}
	return &dto.FindByIdOutput{
		ID:        balance.ID,
		AccountID: balance.AccountID,
		Balance:   balance.Balance,
	}, nil
}
