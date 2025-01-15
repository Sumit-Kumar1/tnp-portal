package migrations

import (
	"gofr.dev/pkg/gofr/migration"
)

const (
	queryCreateUserEnum  = `CREATE TYPE users_role AS ENUM ('admin', 'student', 'maintainer', 'management', 'tnp_unit');`
	queryUserTableCreate = `CREATE TABLE IF NOT EXISTS users(
		user_id VARCHAR(36) PRIMARY KEY,
		name VARCHAR(50) NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL,
		password_hash BYTEA NOT NULL,
		role user_role NOT NULL,
		created_at TIMESTAMPTZ NOT NULL,
		updated_at TIMESTAMPTZ);`
)

func CreateUserTable() migration.Migrate {
	return migration.Migrate{
		UP: func(d migration.Datasource) error {
			if _, err := d.SQL.Exec(queryCreateUserEnum); err != nil {
				return err
			}

			if _, err := d.SQL.Exec(queryUserTableCreate); err != nil {
				return err
			}

			return nil
		},
	}
}
