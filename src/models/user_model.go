package models

import (
	"time"
)

type UserModel struct {
	ID        	uint64		`gorm:"primarykey;autoIncrement" db:"id" json:"id" validate:"required"`
	Email     	string		`gorm:"unique" db:"email" json:"email" validate:"required,string"`
	Password  	string		`db:"password" json:"password" validate:"required,string"`
	CompanyName string 		`db:"company_name" json:"company_name" validate:"required,string"`
	CreatedAt 	time.Time 	`db:"created_at" json:"created_at"`
}