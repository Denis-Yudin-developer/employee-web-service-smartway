package service

import (
	"smartway-test-task/internal/model"
	"smartway-test-task/internal/repository"
)

type Employee interface {
	Create(employee model.Employee) (int, error)
	GetAll() ([]model.Employee, error)
	GetById(id int) (model.Employee, error)
	GetAllByCompany(companyId int) ([]model.EmployeeResponse, error)
	GetAllByDepartment(department string) ([]model.EmployeeResponse, error)
	Update(updatedEmployee model.UpdateEmployee, employeeId int) error
	Delete(id int) error
}

type Service struct {
	Employee
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Employee: NewEmployeeService(repos),
	}
}
