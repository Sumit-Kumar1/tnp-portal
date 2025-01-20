package migrations

import (
	"gofr.dev/pkg/gofr/migration"
)

const (
	queryUserTableCreate = `CREATE TABLE IF NOT EXISTS users(
		user_id UUID PRIMARY KEY,
		enrollment_id VARCHAR(10) NOT NULL,	
		name VARCHAR(50) NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL,
		password_hash BYTEA NOT NULL,
		role VARCHAR(50) NOT NULL,
		created_at TIMESTAMPTZ NOT NULL,
		updated_at TIMESTAMPTZ);`
)

func CreateUserTable() migration.Migrate {
	return migration.Migrate{
		UP: func(d migration.Datasource) error {
			if _, err := d.SQL.Exec(queryUserTableCreate); err != nil {
				return err
			}

			return nil
		},
	}
}
