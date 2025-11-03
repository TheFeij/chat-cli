package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type PostgreSQL struct {
	client *gorm.DB
}

func NewPostgreSQL() *PostgreSQL {
	dsn := os.Getenv("POSTGRES_DATABASE")
	if dsn == "" {
		log.Fatal("POSTGRES_DATABASE environment variable not set")
	}

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err.Error())
	}

	return &PostgreSQL{client: db}
}

func (p *PostgreSQL) Close() error {
	sqlDB, err := p.client.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}
