package dto

type LoginDto struct {
	Email 		string `json:"email" validate:"required,string"`
	Password 	string `json:"password" validate:"required,string"`
}