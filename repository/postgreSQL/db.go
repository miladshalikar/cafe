package postgreSQL

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

type Config struct {
	Host     string `koanf:"host"`
	Port     string `koanf:"port"`
	Username string `koanf:"username"`
	Password string `koanf:"password"`
	DBName   string `koanf:"db_name"`
	SSLMode  string `koanf:"ssl_mode"`
}

type DB struct {
	config Config
	db     *sql.DB
}

func (m DB) Conn() *sql.DB {
	return m.db
}

const (
	dbMaxIdleConns    = 10
	dbMaxOpenConns    = 10
	dbMaxConnLifetime = time.Minute * 3
)

func New(cfg Config) *DB {

	cnn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Tehran",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password,
		cfg.DBName, cfg.SSLMode)

	dbClient, err := sql.Open("postgres", cnn)
	if err != nil {
		panic(fmt.Errorf("can't open mysql db: %w", err))
	}

	dbClient.SetMaxIdleConns(dbMaxIdleConns)
	dbClient.SetMaxOpenConns(dbMaxOpenConns)
	dbClient.SetConnMaxLifetime(dbMaxConnLifetime)

	return &DB{config: cfg, db: dbClient}
}
