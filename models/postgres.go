package models

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib" // Added pgx driver import
)

type Postgres struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

func (cfg Postgres) String() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Database,
		cfg.SSLMode,
	)
}

func DefaultPostgresConfig() Postgres {
	return Postgres{
		Host:     "localhost",
		Port:     "5432",
		User:     "pruthvij",
		Password: "pruthvij008",
		Database: "newone",
		SSLMode:  "disable",
	}
}

func Open(config Postgres) (*sql.DB, error) {
	db, err := sql.Open("pgx", config.String())
	if err != nil {
		return nil, fmt.Errorf("open: %w", err)
	}
	return db, nil
}
