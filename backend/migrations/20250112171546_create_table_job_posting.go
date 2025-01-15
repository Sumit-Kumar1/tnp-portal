package migrations

import "gofr.dev/pkg/gofr/migration"

const (
	queryJobPostingCreateTable = `CREATE TABLE IF NOT EXISTS job_posting(
		job_id UUID PRIMARY KEY,
		employer_name VARCHAR(50) NOT NULL,
		description TEXT NOT NULL,
		location VARCHAR(50) NOT NULL,
		qualifications TEXT NOT NULL,
		salary NUMERIC(10,0),
		doc_id UUID,
		created_at TIMESTAMPTZ NOT NULL,
		updated_at TIMESTAMPTZ);`
)

func CreateJobPostingTable() migration.Migrate {
	return migration.Migrate{
		UP: func(d migration.Datasource) error {
			if _, err := d.SQL.Exec(queryJobPostingCreateTable); err != nil {
				return err
			}

			return nil
		},
	}
}
