package dto

type RegisterDto struct {
	CompanyName string `json:"company_name" validate:"required,string"`
	Email 		string `json:"email" validate:"required,string"`
	Password 	string `json:"password" validate:"required,string"`
}