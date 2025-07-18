package usecase

import (
	"microservice/domain/dto"
	"microservice/domain/entity"
	"microservice/domain/interfaces"
)

type CreateBalanceUseCase struct {
	BalanceRepository interfaces.IBalanceRepository
}

func NewCreateBalanceUseCase(balanceRepository interfaces.IBalanceRepository) *CreateBalanceUseCase {
	return &CreateBalanceUseCase{
		BalanceRepository: balanceRepository,
	}
}

func (c *CreateBalanceUseCase) Execute(input *dto.CreateBalanceInput) (*entity.Balance, error) {
	balance := entity.NewBalance(input.AccountID, input.Balance)
	err := c.BalanceRepository.Save(balance)
	if err != nil {
		return nil, err
	}
	return balance, nil
}
