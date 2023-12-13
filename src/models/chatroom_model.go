package models

import (
	"time"
)

type ChatRoomModel struct {
	ID       uint64		`gorm:"primarykey;autoIncrement" db:"id" json:"id" validate:"required"`
	Name     string		`db:"name" json:"name" validate:"required,string"`
	UserID   uint64 	`db:"user_id" json:"user_id" validate:"required"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}