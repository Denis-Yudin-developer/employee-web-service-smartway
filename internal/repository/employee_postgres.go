package repository

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"smartway-test-task/internal/model"
)

type EmployeePostgres struct {
	pg *PostgresDb
}

func (r *EmployeePostgres) Create(employee model.Employee) (int, error) {
	tx, err := r.pg.db.Begin()
	if err != nil {
		return 0, err
	}

	var employeeId int
	row, err := r.pg.dot.QueryRow(tx, createUser,
		employee.Name,
		employee.Surname,
		employee.Phone,
		employee.CompanyId)

	if err != nil {
		return 0, err
	}

	err = row.Scan(&employeeId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return employeeId, tx.Commit()
}

func (r *EmployeePostgres) GetAll() ([]model.Employee, error) {
	//TODO implement me
	panic("implement me")
}

func (r *EmployeePostgres) GetById(id int) (model.Employee, error) {
	//TODO implement me
	panic("implement me")
}

func (r *EmployeePostgres) GetAllByCompany(companyId int) ([]model.EmployeeResponse, error) {
	employees := make([]model.EmployeeResponse, 0)

	rows, err := r.pg.dot.Query(r.pg.db, findEmployeesByCompanyId, companyId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var employeeResponse model.EmployeeResponse
		var departmentResponse model.DepartmentResponse
		var passportResponse model.PassportResponse
		employeeResponse.Department = &departmentResponse
		employeeResponse.Passport = &passportResponse

		err = rows.Scan(&employeeResponse.Id,
			&employeeResponse.Name,
			&employeeResponse.Surname,
			&employeeResponse.Phone,
			&employeeResponse.CompanyId,
			&employeeResponse.Department.Id,
			&employeeResponse.Department.Name,
			&employeeResponse.Department.Phone,
			&employeeResponse.Passport.Id,
			&employeeResponse.Passport.PassportType,
			&employeeResponse.Passport.PassportNumber)
		if err != nil {
			return nil, err
		}

		employees = append(employees, employeeResponse)
	}
	return employees, nil
}

func (r *EmployeePostgres) GetAllByDepartment(department string) ([]model.EmployeeResponse, error) {
	employees := make([]model.EmployeeResponse, 0)

	rows, err := r.pg.dot.Query(r.pg.db, findEmployeesByDepartment, department)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var employeeResponse model.EmployeeResponse
		var departmentResponse model.DepartmentResponse
		var passportResponse model.PassportResponse
		employeeResponse.Department = &departmentResponse
		employeeResponse.Passport = &passportResponse

		err = rows.Scan(&employeeResponse.Id,
			&employeeResponse.Name,
			&employeeResponse.Surname,
			&employeeResponse.Phone,
			&employeeResponse.CompanyId,
			&employeeResponse.Department.Id,
			&employeeResponse.Department.Name,
			&employeeResponse.Department.Phone,
			&employeeResponse.Passport.Id,
			&employeeResponse.Passport.PassportType,
			&employeeResponse.Passport.PassportNumber)

		if err != nil {
			return nil, err
		}

		employees = append(employees, employeeResponse)
	}
	return employees, nil
}

func (r *EmployeePostgres) IsEmployeePresent(employeeId int) bool {
	row, err := r.pg.dot.QueryRow(r.pg.db, isEmployeePresent, employeeId)
	if err != nil {
		logrus.Print(err.Error())
		return false
	}
	if err := row.Scan(&employeeId); err != nil {
		if err != sql.ErrNoRows {
			logrus.Print(err.Error())
		}
		return false
	}
	return true
}

func (r *EmployeePostgres) Update(updatedEmployee model.UpdateEmployee, employeeId int) error {
	var result []interface{}

	result = append(result, updatedEmployee.Name,
		updatedEmployee.Surname,
		updatedEmployee.Phone,
		updatedEmployee.CompanyId,
		employeeId)

	_, err := r.pg.dot.Exec(r.pg.db, updateEmployee, result...)

	if err != nil {
		return err
	}
	return nil
}

func (r *EmployeePostgres) Delete(id int) error {
	_, err := r.pg.dot.Exec(r.pg.db, deleteUser, id)

	return err
}

func NewEmployeePostgres(db *PostgresDb) *EmployeePostgres {
	return &EmployeePostgres{pg: db}
}
