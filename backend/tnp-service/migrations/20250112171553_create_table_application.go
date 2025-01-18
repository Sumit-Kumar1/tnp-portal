package migrations

import "gofr.dev/pkg/gofr/migration"

const (
	queryApplicationTableCreate = `CREATE TYPE application_status AS ENUM('applied', 'shortlisted', 'rejected', 'pending', 'withdraw');
	CREATE TABLE IF NOT EXISTS application(
		application_id UUID PRIMARY KEY,
		job_id UUID NOT NULL,
		student_id UUID NOT NULL,
		status application_status,
		applied_at TIMESTAMPTZ NOT NULL,
		updated_at TIMESTAMPTZ);`
)

func CreateApplicationsTable() migration.Migrate {
	return migration.Migrate{
		UP: func(d migration.Datasource) error {
			if _, err := d.SQL.Exec(queryApplicationTableCreate); err != nil {
				return err
			}

			return nil
		},
	}
}
