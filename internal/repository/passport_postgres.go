package repository

import "smartway-test-task/internal/model"

type PassportPostgres struct {
	pg *PostgresDb
}

func (r *PassportPostgres) Create(passport model.Passport, employeeId int) error {
	tx, err := r.pg.db.Begin()
	if err != nil {
		return err
	}

	var passportId int
	row, err := r.pg.dot.QueryRow(tx, createPassport,
		employeeId,
		passport.PassportType,
		passport.PassportNumber)

	if err != nil {
		return err
	}

	err = row.Scan(&passportId)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = r.pg.dot.Exec(tx, addPassportToUser,
		passportId,
		employeeId)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *PassportPostgres) Update(updatedPassport model.UpdatePassport, employeeId int) error {
	var result []interface{}

	result = append(result, updatedPassport.PassportType)
	result = append(result, updatedPassport.PassportNumber)
	result = append(result, employeeId)

	_, err := r.pg.dot.Exec(r.pg.db, updatePassportByUserId, result...)
	if err != nil {
		return err
	}
	return nil
}

func NewPassportPostgres(db *PostgresDb) *PassportPostgres {
	return &PassportPostgres{pg: db}
}
