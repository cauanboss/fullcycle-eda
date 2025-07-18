package interfaces

type IBalanceUseCase interface {
	Execute(input any) (any, error)
}
