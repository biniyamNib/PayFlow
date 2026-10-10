package repository

import "database/sql"

type DBExecutor interface {
	QueryRow(query string, args ...any) *sql.Row
}