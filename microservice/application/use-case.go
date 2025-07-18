package application

import (
	"microservice/domain/dto"
	"microservice/domain/entity"
	"microservice/domain/interfaces"
)

type BalanceUseCases struct {
	CreateBalance interfaces.IBalanceUseCase[*entity.Balance, *dto.CreateBalanceInput]
	FindById      interfaces.IBalanceUseCase[*dto.FindByIdOutput, *dto.FindByIdInput]
	Find          interfaces.IBalanceUseCase[[]*entity.Balance, any]
}

func NewBalanceUseCases(
	createBalance interfaces.IBalanceUseCase[*entity.Balance, *dto.CreateBalanceInput],
	findByIdUseCase interfaces.IBalanceUseCase[*dto.FindByIdOutput, *dto.FindByIdInput],
	findUseCase interfaces.IBalanceUseCase[[]*entity.Balance, any],
) *BalanceUseCases {
	return &BalanceUseCases{
		CreateBalance: createBalance,
		FindById:      findByIdUseCase,
		Find:          findUseCase,
	}
}
