package initiator

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func InitMigiration(path string, db_url string) *migrate.Migrate {

	// absPath, err := filepath.Abs(path)
	// if err != nil {
	// 	log.Fatal("Error getting absolute path:", err)
	// }

	m_file, err := migrate.New(fmt.Sprintf("file://%s", path), db_url)

	if err != nil {

		log.Fatal(err)
	}

	return m_file

}

func UpMigiration(m *migrate.Migrate) {

	fmt.Println("===========Migrate Up====================")

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

}
