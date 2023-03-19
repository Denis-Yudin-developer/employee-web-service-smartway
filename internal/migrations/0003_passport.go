package migrations

import (
	"database/sql"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upPassport, downPassport)
}

func upPassport(tx *sql.Tx) error {
	query := `CREATE TABLE IF NOT EXISTS passport (
  				"id" SERIAL PRIMARY KEY, 
  				"employee_id" INTEGER REFERENCES employees(id) ON DELETE CASCADE,	
                "passport_type" VARCHAR(50),
                "passport_number" VARCHAR(15));`
	_, err := tx.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func downPassport(tx *sql.Tx) error {
	query := `DROP TABLE passport;`
	_, err := tx.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
