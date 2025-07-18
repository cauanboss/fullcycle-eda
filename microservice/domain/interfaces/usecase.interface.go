package interfaces

type IBalanceUseCase[T any, I any] interface {
	Execute(input I) (T, error)
}
