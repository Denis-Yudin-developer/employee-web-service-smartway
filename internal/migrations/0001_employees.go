package migrations

import (
	"database/sql"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upEmployees, downEmployees)
}

func upEmployees(tx *sql.Tx) error {
	query := `CREATE TABLE IF NOT EXISTS employees (
  				"id" SERIAL PRIMARY KEY, 
                "name" VARCHAR(50),
    			"surname" VARCHAR(50),
                "phone" VARCHAR(15),
    			"company_id" int,
    			"department_id" int,
    			"passport_id" int);`
	_, err := tx.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func downEmployees(tx *sql.Tx) error {
	query := `DROP TABLE employees;`
	_, err := tx.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
