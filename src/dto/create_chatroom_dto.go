package dto

type CreateChatRoomDto struct {
	Name 		string `json:"name" validate:"required,string"`
}