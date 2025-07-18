package interfaces

import "database/sql"

type IDatabaseAdapter interface {
	GetDB() *sql.DB
	Close() error
}
