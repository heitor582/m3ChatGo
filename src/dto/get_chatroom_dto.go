package dto

import (
	"time"
	models"github.com/heitor582/m3ChatGo/src/models"
)

type GetChatRoomDto struct {
	ID       uint64		`json:"id" validate:"required"`
	Name     string		`json:"name" validate:"required,string"`
	Messages []GetChatRoomMessage `json:"messages"`
	CreatedAt time.Time `json:"created_at"`
}

type GetChatRoomMessage struct {
	models.MessageModel
	CompanyName 	string 		`json:"company_name" validate:"required,string"`
}