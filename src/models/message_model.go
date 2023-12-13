package models

import "time"

type MessageModel struct {
	ID        		uint64		`gorm:"primarykey;autoIncrement" db:"id" json:"id" validate:"required"`
	UserID   		uint64 		`db:"user_id" json:"user_id" validate:"required"`
	ChatRoomId     	uint64		`db:"chatroom_id" json:"chatroom_id" validate:"required"`
	Content     	string		`db:"content" json:"content" validate:"required, string"`
	CreatedAt	 	time.Time 	`db:"created_at" json:"created_at"`
}