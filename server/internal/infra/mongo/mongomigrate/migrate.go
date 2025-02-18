package mongomigrate

import "go.mongodb.org/mongo-driver/mongo"

type Migrator struct {
	db *mongo.Database
}

func NewMigrator(db *mongo.Database) *Migrator {
	return &Migrator{db: db}
}

func (m *Migrator) Migrate1() error {
	// TODO: implement
	return nil
}
