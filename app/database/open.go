package database

import (
	"fmt"

	// load postgres driver
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

// Parameters used to start a database connection
type DBParams struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
	Database string `json:"database,omitempty"`
	Schema   string `json:"schema,omitempty"`
	Debug    bool   `json:"debug,omitempty"`
}

// Opens a database connection pool
func OpenDatabase(params *DBParams) (*Service, error) {
	host := params.Host
	if host == "" {
		host = "localhost"
	}
	port := params.Port
	if port == 0 {
		port = 5432
	}

	template := "postgres://%s:%s@%s:%d/%s"
	url := fmt.Sprintf(template, params.Username, params.Password, host, port, params.Database)

	db, err := sqlx.Open("pgx", url)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	svc := NewService(params.Debug, db)

	return svc, nil
}
