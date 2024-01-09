package dto

import "time"

type MessageDto struct {
	CompanyName 	string 		`json:"company_name" validate:"string"`
	UserID   		uint64 		`json:"user_id"`
	CreatedAt 	time.Time 		`json:"created_at"`
	ChatRoomId     	uint64		`json:"chatroom_id" validate:"required"`
	Content			string 		`json:"content" validate:"required,string"`
}