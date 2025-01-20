package migrations

import "gofr.dev/pkg/gofr/migration"

const (
	queryStudentTableCreate = `CREATE TABLE IF NOT EXISTS student(
		student_id UUID PRIMARY KEY,
		user_id UUID NOT NULL,
		name VARCHAR(50) NOT NULL,
		degree VARCHAR(50) NOT NULL,
		department VARCHAR(50) NOT NULL,
		major VARCHAR(50) NOT NULL,
		grad_year INTEGER NOT NULL,
		contact_info VARCHAR(10) NOT NULL,
		educational_background TEXT ARRAY,
		skills TEXT ARRAY,
		doc_id UUID,
		created_at TIMESTAMPTZ NOT NULL,
		updated_at TIMESTAMPTZ);`
)

func CreateStudentTable() migration.Migrate {
	return migration.Migrate{
		UP: func(d migration.Datasource) error {
			if _, err := d.SQL.Exec(queryStudentTableCreate); err != nil {
				return err
			}

			return nil
		},
	}
}
