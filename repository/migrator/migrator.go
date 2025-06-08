package migrator

import (
	"database/sql"
	"fmt"
	"github.com/miladshalikar/cafe/repository/postgreSQL"
	migrate "github.com/rubenv/sql-migrate"
)

type Migrator struct {
	dialect    string
	dbConfig   postgreSQL.Config
	migrations *migrate.FileMigrationSource
}

func New(dbConfig postgreSQL.Config) Migrator {
	migrations := &migrate.FileMigrationSource{Dir: "repository/postgreSQL/migration"}
	return Migrator{dbConfig: dbConfig, migrations: migrations, dialect: "postgres"}
}

func (m Migrator) Up() {
	cnn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Tehran",
		m.dbConfig.Host, m.dbConfig.Port, m.dbConfig.Username, m.dbConfig.Password, m.dbConfig.DBName, m.dbConfig.SSLMode)

	fmt.Printf("postgres add= %+v\n", cnn)

	db, err := sql.Open(m.dialect, cnn)
	if err != nil {
		panic(fmt.Errorf("can't open postgres db: %w", err))
	}

	n, err := migrate.Exec(db, m.dialect, m.migrations, migrate.Up)
	if err != nil {
		panic(fmt.Errorf("can't apply migrations: %w", err))
	}
	fmt.Printf("Applied %d migrations!\n", n)
}

func (m Migrator) Down() {
	cnn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Tehran",
		m.dbConfig.Host, m.dbConfig.Port, m.dbConfig.Username, m.dbConfig.Password, m.dbConfig.DBName, m.dbConfig.SSLMode)

	db, err := sql.Open(m.dialect, cnn)
	if err != nil {
		panic(fmt.Errorf("can't open postgres db: %w", err))
	}

	n, err := migrate.Exec(db, m.dialect, m.migrations, migrate.Down)
	if err != nil {
		panic(fmt.Errorf("can't rollback migrations: %w", err))
	}
	fmt.Printf("Rollbacked %d migrations!\n", n)
}
