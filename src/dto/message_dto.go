package dto

type MessageDto struct {
	CompanyName 	string 		`json:"company_name" validate:"required,string"`
	ChatRoomId     	uint64		`json:"chatroom_id" validate:"required"`
	Content			string 		`json:"content" validate:"required,string"`
}