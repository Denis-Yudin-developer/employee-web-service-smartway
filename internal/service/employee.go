package service

import (
	"fmt"
	"smartway-test-task/internal/model"
	"smartway-test-task/internal/repository"
)

type EmployeeService struct {
	employeeRepository   repository.Employee
	departmentRepository repository.Department
	passportRepository   repository.Passport
}

func (e *EmployeeService) Create(employee model.Employee) (int, error) {
	createdEmployeeId, err := e.employeeRepository.Create(employee)
	if err != nil {
		return 0, err
	}

	department := employee.Department
	err = e.departmentRepository.Create(*department, createdEmployeeId)
	if err != nil {
		return 0, err
	}

	passport := *employee.Passport
	err = e.passportRepository.Create(passport, createdEmployeeId)
	if err != nil {
		return 0, err
	}

	return createdEmployeeId, nil
}

func (e *EmployeeService) GetAll() ([]model.Employee, error) {
	panic("implement me")
}

func (e *EmployeeService) GetById(id int) (model.Employee, error) {
	panic("implement me")
}

func (e *EmployeeService) GetAllByCompany(companyId int) ([]model.EmployeeResponse, error) {
	employees, err := e.employeeRepository.GetAllByCompany(companyId)
	if err != nil {
		return nil, err
	}
	return employees, nil
}

func (e *EmployeeService) GetAllByDepartment(department string) ([]model.EmployeeResponse, error) {
	employees, err := e.employeeRepository.GetAllByDepartment(department)
	if err != nil {
		return nil, err
	}
	return employees, nil
}

func (e *EmployeeService) Update(updatedEmployee model.UpdateEmployee, employeeId int) error {
	isPresent := e.employeeRepository.IsEmployeePresent(employeeId)
	if !isPresent {
		return fmt.Errorf("ОШИБКА ОБНОВЛЕНИЯ. РАБОТНИК С ID %d ОТСУТСТВУЕТ", employeeId)
	}
	err := e.employeeRepository.Update(updatedEmployee, employeeId)
	if err != nil {
		return err
	}
	updatedEmployeeDepartment := updatedEmployee.Department
	if updatedEmployeeDepartment == nil {
		return nil
	}
	err = e.departmentRepository.Update(*updatedEmployeeDepartment, employeeId)
	if err != nil {
		return err
	}
	updatedEmployeePassport := updatedEmployee.Passport
	if updatedEmployeePassport == nil {
		return nil
	}
	err = e.passportRepository.Update(*updatedEmployeePassport, employeeId)
	if err != nil {
		return err
	}
	return nil
}

func (e *EmployeeService) Delete(employeeId int) error {
	isPresent := e.employeeRepository.IsEmployeePresent(employeeId)
	if !isPresent {
		return fmt.Errorf("ОШИБКА УДАЛЕНИЯ. РАБОТНИК С ID %d ОТСУТСТВУЕТ", employeeId)
	}
	err := e.employeeRepository.Delete(employeeId)
	return err
}

func NewEmployeeService(repo *repository.Repository) *EmployeeService {
	return &EmployeeService{
		employeeRepository:   repo.Employee,
		departmentRepository: repo.Department,
		passportRepository:   repo.Passport,
	}
}
