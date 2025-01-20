package migrations

import "gofr.dev/pkg/gofr/migration"

const (
	queryDocumentTableCreate = `CREATE TABLE IF NOT EXISTS document(
		doc_id UUID PRIMARY KEY,
		user_id UUID NOT NULL,
		doc_type VARCHAR(10) NOT NULL,
		document bytea NOT NULL,
		created_at TIMESTAMPTZ NOT NULL,
		updated_at TIMESTAMPTZ);`
)

func CreateDocumentsTable() migration.Migrate {
	return migration.Migrate{
		UP: func(d migration.Datasource) error {
			if _, err := d.SQL.Exec(queryDocumentTableCreate); err != nil {
				return err
			}

			return nil
		},
	}
}
