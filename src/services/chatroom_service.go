package services

import (
	"errors"
	"time"

	"github.com/heitor582/m3ChatGo/src/configuration"
	dto "github.com/heitor582/m3ChatGo/src/dto"
	models "github.com/heitor582/m3ChatGo/src/models"
)

func GetAllChatRooms(userId uint64) []models.ChatRoomModel {
	db := configuration.DBConn
	var chatRooms []models.ChatRoomModel
	db.Find(&chatRooms, "user_id = ?", userId)
	return chatRooms
}

func GetChatRoom(chatRoomId uint64, userId uint64) (dto.GetChatRoomDto, error) {
	db := configuration.DBConn
	var chatRoom models.ChatRoomModel
	db.Where("id = ? AND user_id = ?", chatRoomId, userId).Find(&chatRoom)
	if chatRoom.ID == 0 {
		return dto.GetChatRoomDto{}, errors.New("ChatRoom was not found")
	}
	var messages []models.MessageModel
	err := db.Where("chat_room_id = ? AND user_id = ?", chatRoomId, userId).Find(&messages).Error
	if err != nil {
		return dto.GetChatRoomDto{}, errors.New(err.Error())
	}
	return dto.GetChatRoomDto{
		ID: chatRoom.ID,
		Name: chatRoom.Name,
		Messages: messages,
		CreatedAt: chatRoom.CreatedAt,
	}, nil
}

func NewChatRoom(createDto dto.CreateChatRoomDto, userId uint64) (models.ChatRoomModel, error) {
	db := configuration.DBConn
	var chatRoom models.ChatRoomModel = models.ChatRoomModel{
		Name: createDto.Name,
		UserID: userId,
		CreatedAt: time.Now(),
	}
	
	err := db.Create(&chatRoom).Error
	if err != nil {
		return models.ChatRoomModel{}, errors.New(err.Error())
	}

	return chatRoom, nil
}

func DeleteChatRoom(id uint64) error {
	db := configuration.DBConn
	var chatRoom models.ChatRoomModel
	db.First(&chatRoom, id)
	if chatRoom.ID == 0 {
		return errors.New("chatroom was not found")
	}
	db.Delete(&chatRoom)
	return nil
}