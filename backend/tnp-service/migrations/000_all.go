package migrations

import (
	"gofr.dev/pkg/gofr/migration"
)

func All() map[int64]migration.Migrate {
	return map[int64]migration.Migrate{
		20250112171512: CreateUserTable(),
		20250112171534: CreateStudentTable(),
		20250112171546: CreateJobPostingTable(),
		20250112171553: CreateApplicationsTable(),
		20250112171559: CreateDocumentsTable(),
	}
}
