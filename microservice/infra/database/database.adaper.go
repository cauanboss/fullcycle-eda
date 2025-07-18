package database

import (
	"database/sql"
	"fmt"
	"microservice/domain/interfaces"
)

type DatabaseAdapter struct {
	db *sql.DB
}

func NewDatabaseAdapter() interfaces.IDatabaseAdapter {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "mysql", "3306", "wallet"))
	if err != nil {
		panic(err)
	}
	return &DatabaseAdapter{db: db}
}

func (d *DatabaseAdapter) GetDB() *sql.DB {
	return d.db
}

func (d *DatabaseAdapter) Close() error {
	return d.db.Close()
}
