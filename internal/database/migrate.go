package database

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
)

// Migrate applies database migrations from the specified path
func Migrate(db *mongo.Client, migrationsPath string) error {
	// Create a MongoDB driver instance
	driver, err := mongodb.WithInstance(db, &mongodb.Config{})
	if err != nil {
		return fmt.Errorf("migration driver error: %w", err)
	}

	// Initialize the migration instance
	m, err := migrate.NewWithDatabaseInstance(
		"file://"+migrationsPath,
		"mongodb", driver,
	)
	if err != nil {
		return fmt.Errorf("migration initialization error: %w", err)
	}

	// Apply migrations
	err = m.Up()
	if err != nil {
		if err == migrate.ErrNoChange {
			log.Println("⚠️ No new migrations to apply")
			return nil
		}
		return fmt.Errorf("migration up error: %w", err)
	}

	log.Println("✅ Database migrated successfully")
	return nil
}

// Rollback rolls back the last applied migration
func Rollback(db *mongo.Client, migrationsPath string) error {
	// Create a MongoDB driver instance
	driver, err := mongodb.WithInstance(db, &mongodb.Config{})
	if err != nil {
		return fmt.Errorf("migration driver error: %w", err)
	}

	// Initialize the migration instance
	m, err := migrate.NewWithDatabaseInstance(
		"file://"+migrationsPath,
		"mongodb", driver,
	)
	if err != nil {
		return fmt.Errorf("migration initialization error: %w", err)
	}

	// Rollback the last migration
	err = m.Down()
	if err != nil {
		if err == migrate.ErrNoChange {
			log.Println("⚠️ No migrations to roll back")
			return nil
		}
		return fmt.Errorf("migration down error: %w", err)
	}

	log.Println("✅ Last migration rolled back successfully")
	return nil
}
