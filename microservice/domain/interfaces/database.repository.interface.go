package interfaces

import "microservice/domain/entity"

type IBalanceRepository interface {
	FindById(id string) (*entity.Balance, error)
	Find() ([]*entity.Balance, error)
	Update(balance *entity.Balance) error
	Save(balance *entity.Balance) error
}
