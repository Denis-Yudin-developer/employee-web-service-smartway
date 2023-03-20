package model

type Department struct {
	Id    int    `json:"-"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type UpdateDepartment struct {
	Name  *string `json:"name"`
	Phone *string `json:"phone"`
}

type DepartmentResponse struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}
