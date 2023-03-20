package model

import "fmt"

type Employee struct {
	Id         int         `json:"-"`
	Name       string      `json:"name"`
	Surname    string      `json:"surname"`
	Phone      string      `json:"phone"`
	CompanyId  int         `json:"company_id"`
	Passport   *Passport   `json:"passport"`
	Department *Department `json:"department"`
}

type UpdateEmployee struct {
	Name       *string           `json:"name"`
	Surname    *string           `json:"surname"`
	Phone      *string           `json:"phone"`
	CompanyId  *int              `json:"company_id"`
	Passport   *UpdatePassport   `json:"passport"`
	Department *UpdateDepartment `json:"department"`
}

type EmployeeResponse struct {
	Id         int                 `json:"id"`
	Name       string              `json:"name"`
	Surname    string              `json:"surname"`
	Phone      string              `json:"phone"`
	CompanyId  int                 `json:"company_id"`
	Passport   *PassportResponse   `json:"passport"`
	Department *DepartmentResponse `json:"department"`
}

func Validate(e Employee) error {
	if e.Name == "" {
		return fmt.Errorf("НЕОБХОДИМО УКАЗАТЬ ИМЯ ПРИ СОЗДАНИИ РАБОТНИКА")
	}
	if e.Surname == "" {
		return fmt.Errorf("НЕОБХОДИМО УКАЗАТЬ ФАМИЛИЮ ПРИ СОЗДАНИИ РАБОТНИКА")
	}

	if e.Phone == "" {
		return fmt.Errorf("НЕОБХОДИМО УКАЗАТЬ ТЕЛЕФОН ПРИ СОЗДАНИИ РАБОТНИКА")
	}
	if e.CompanyId == 0 {
		return fmt.Errorf("НЕОБХОДИМО УКАЗАТЬ НОМЕР КОМПАНИИ РАБОТНИКА")
	}

	if e.Department == nil || e.Department.Name == "" || e.Department.Phone == "" {
		return fmt.Errorf("НЕОБХОДИМО УКАЗАТЬ ДАННЫЕ О ДЕПАРТАМЕНТЕ ПРИ СОЗДАНИИ РАБОТНИКА")
	}
	if e.Passport == nil || e.Passport.PassportType == "" || e.Passport.PassportNumber == "" {
		return fmt.Errorf("НЕОБХОДИМО УКАЗАТЬ ПАСПОРТНЫЕ ДАННЫЕ ПРИ СОЗДАНИИ РАБОТНИКА")
	}
	return nil
}
