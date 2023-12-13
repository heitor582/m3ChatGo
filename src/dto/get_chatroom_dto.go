package dto

import (
	"time"
	models"github.com/heitor582/m3ChatGo/src/models"
)

type GetChatRoomDto struct {
	ID       uint64		`json:"id" validate:"required"`
	Name     string		`json:"name" validate:"required,string"`
	Messages []models.MessageModel `json:"messages"`
	CreatedAt time.Time `json:"created_at"`
}