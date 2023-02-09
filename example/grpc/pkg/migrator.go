package pkg

import (
	gkit "gkit"
	"log"
)

type Migrator struct {
}

func NewMigrator() gkit.Migrator {
	return &Migrator{}
}

func (m *Migrator) Migrate() error {
	log.Println("Starting migration ...")
	log.Println("Migration done ...")
	return nil
}
