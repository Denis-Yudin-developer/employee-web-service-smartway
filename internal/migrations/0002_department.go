package migrations

import (
	"database/sql"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upDepartment, downDepartment)
}

func upDepartment(tx *sql.Tx) error {
	query := `CREATE TABLE IF NOT EXISTS department (
  				"id" SERIAL PRIMARY KEY, 
  				"employee_id" INTEGER REFERENCES employees(id) ON DELETE CASCADE,
                "name" VARCHAR(50),
                "phone" VARCHAR(15));`
	_, err := tx.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func downDepartment(tx *sql.Tx) error {
	query := `DROP TABLE department;`
	_, err := tx.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
