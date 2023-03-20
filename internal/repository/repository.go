package repository

import (
	"smartway-test-task/internal/model"
)

type Employee interface {
	Create(employee model.Employee) (int, error)
	GetAll() ([]model.Employee, error)
	GetById(id int) (model.Employee, error)
	GetAllByCompany(companyId int) ([]model.EmployeeResponse, error)
	GetAllByDepartment(department string) ([]model.EmployeeResponse, error)
	IsEmployeePresent(employeeId int) bool
	Update(updatedEmployee model.UpdateEmployee, employeeId int) error
	Delete(id int) error
}

type Department interface {
	Create(department model.Department, employeeId int) error
	Update(updatedDepartment model.UpdateDepartment, employeeId int) error
}

type Passport interface {
	Create(passport model.Passport, employeeId int) error
	Update(updatedPassport model.UpdatePassport, employeeId int) error
}

type Repository struct {
	Employee
	Department
	Passport
}

func NewRepository(pg *PostgresDb) *Repository {
	return &Repository{
		Employee:   NewEmployeePostgres(pg),
		Department: NewDepartmentPostgres(pg),
		Passport:   NewPassportPostgres(pg),
	}
}
