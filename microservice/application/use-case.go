package application

import "microservice/domain/interfaces"

type BalanceUseCases struct {
	CreateBalance interfaces.IBalanceUseCase
	FindById      interfaces.IBalanceUseCase
	Find          interfaces.IBalanceUseCase
}

func NewBalanceUseCases(createBalance, findById, find interfaces.IBalanceUseCase) *BalanceUseCases {
	return &BalanceUseCases{
		CreateBalance: createBalance,
		FindById:      findById,
		Find:          find,
	}
}
