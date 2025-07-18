package usecase

import (
	"microservice/domain/entity"
	"microservice/domain/interfaces"
)

type FindUseCase struct {
	BalanceRepository interfaces.IBalanceRepository
}

func NewFindUseCase(balanceRepository interfaces.IBalanceRepository) *FindUseCase {
	return &FindUseCase{
		BalanceRepository: balanceRepository,
	}
}

func (f *FindUseCase) Execute(input any) ([]*entity.Balance, error) {
	balance, err := f.BalanceRepository.Find()
	if err != nil {
		return nil, err
	}
	return balance, nil
}
