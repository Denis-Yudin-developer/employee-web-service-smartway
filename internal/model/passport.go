package model

type Passport struct {
	Id             int    `json:"-"`
	PassportType   string `json:"passport_type"`
	PassportNumber string `json:"passport_number"`
}

type UpdatePassport struct {
	PassportType   *string `json:"passport_type"`
	PassportNumber *string `json:"passport_number"`
}

type PassportResponse struct {
	Id             int    `json:"id"`
	PassportType   string `json:"passport_type"`
	PassportNumber string `json:"passport_number"`
}
