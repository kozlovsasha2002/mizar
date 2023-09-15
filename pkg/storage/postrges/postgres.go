package postrges

import (
	"database/sql"
	"fmt"
	"mizar/internal/config"
)

type Storage struct {
	db *sql.DB
}

func New(cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DbName, cfg.Sslmode))

	op := "storage.postgres.New"
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return db, nil
}
