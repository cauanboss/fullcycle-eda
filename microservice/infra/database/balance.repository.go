package database

import (
	"database/sql"
	"fmt"
	"microservice/domain/entity"
	"microservice/domain/interfaces"
)

type BalanceRepository struct {
	db *sql.DB
}

func NewBalanceRepository(db *sql.DB) interfaces.IBalanceRepository {
	return &BalanceRepository{
		db: db,
	}
}

func (r *BalanceRepository) FindById(id string) (*entity.Balance, error) {
	stmt, err := r.db.Prepare("SELECT id, account_id, balance, created_at FROM balances WHERE id = ?")
	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(id)
	var balance entity.Balance
	err = row.Scan(&balance.ID, &balance.AccountID, &balance.Balance, &balance.CreatedAt)
	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}
	return &balance, nil
}

func (r *BalanceRepository) Save(balance *entity.Balance) error {
	stmt, err := r.db.Prepare("INSERT INTO balances (id, account_id, balance, created_at) VALUES (?, ?, ?, ?)")
	if err != nil {
		fmt.Println("err", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(balance.ID, balance.AccountID, balance.Balance, balance.CreatedAt)
	if err != nil {
		fmt.Println("err", err)
		return err
	}
	return nil
}

func (r *BalanceRepository) Update(balance *entity.Balance) error {
	stmt, err := r.db.Prepare("UPDATE balances SET balance = ? WHERE id = ?")
	if err != nil {
		fmt.Println("err", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(balance.Balance, balance.ID)
	if err != nil {
		fmt.Println("err", err)
		return err
	}
	return nil
}

func (r *BalanceRepository) Find() ([]*entity.Balance, error) {
	stmt, err := r.db.Prepare("SELECT id, account_id, balance, created_at FROM balances")
	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}
	defer rows.Close()

	var balances []*entity.Balance
	for rows.Next() {
		var balance entity.Balance
		err = rows.Scan(&balance.ID, &balance.AccountID, &balance.Balance, &balance.CreatedAt)
		if err != nil {
			fmt.Println("err", err)
			return nil, err
		}
		balances = append(balances, &balance)
	}
	return balances, nil
}

// Verificar se implementa a interface
var _ interfaces.IBalanceRepository = (*BalanceRepository)(nil)
