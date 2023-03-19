package repository

import "smartway-test-task/internal/model"

type DepartmentPostgres struct {
	pg *PostgresDb
}

func (r *DepartmentPostgres) Create(department model.Department, employeeId int) error {
	tx, err := r.pg.db.Begin()
	if err != nil {
		return err
	}

	var departmentId int
	row, err := r.pg.dot.QueryRow(tx, createDepartment,
		employeeId,
		department.Name,
		department.Phone)

	if err != nil {
		return err
	}

	err = row.Scan(&departmentId)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = r.pg.dot.Exec(tx, addDepartmentToUser,
		departmentId,
		employeeId)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *DepartmentPostgres) Update(updatedDepartment model.UpdateDepartment, employeeId int) error {
	var result []interface{}

	result = append(result, updatedDepartment.Name, updatedDepartment.Phone, employeeId)

	_, err := r.pg.dot.Exec(r.pg.db, updateDepartmentByUserId, result...)
	if err != nil {
		return err
	}
	return nil
}

func NewDepartmentPostgres(db *PostgresDb) *DepartmentPostgres {
	return &DepartmentPostgres{pg: db}
}
