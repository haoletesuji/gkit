package gkit

type SQLMigrator struct {
	migrator Migrator
}

func NewSQLMigrator(migrator Migrator) *SQLMigrator {
	return &SQLMigrator{
		migrator,
	}
}

func (m *SQLMigrator) Migrate() error {
	return m.migrator.Migrate()
}
