package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/heitor582/m3ChatGo/src/configuration"
	dto "github.com/heitor582/m3ChatGo/src/dto"
	models "github.com/heitor582/m3ChatGo/src/models"
)

func SendMessageToChatGpt(userMessage string, chatRoomId uint64, userId uint64) (models.MessageModel, error) {
	db := configuration.DBConn
	url := "https://api.openai.com/v1/chat/completions"

	bearerToken := os.Getenv("CHAT_GPT_TOKEN")

	var oldMessages []models.MessageModel
	err := db.Where("user_id IN ?", []uint64{1, userId}).Find(&oldMessages).Error
	if err != nil {
		return models.MessageModel{}, errors.New(err.Error())
	}

	var oldMessagesFormated []dto.ChatGptMessage

	for i := 0; i < len(oldMessages); i++ {
		message := oldMessages[i]
		var role string 

		if message.UserID == 1 {
			role = "system"
		} else {
			role = "user"
		}

		oldMessagesFormated = append(oldMessagesFormated, dto.ChatGptMessage{
			Role: role,
			Content: message.Content,
		})
	}
	
	messages := []dto.ChatGptMessage{
			{
				Role: "system",
				Content: `Você é um consultor empresarial experiente, que trabalha com foco em organizar e crescer empresas.
				Você usa o idioma Português Brasileiro e tem uma comunicação clara e objetiva. Além de estar em uma posição que guia a conversa para a organização e crescimento constante da empresas.
				Seu fluxo de trabalho é primeiramente entender a empresa do cliente ao completo e depois orienta-lo. 
				Primeiro peça informações sobre sua empresa como os tópicos do Business Model Canva. (Detalhe o que é cada tópicos)
				Segundo analise os pontos que um consultor empresarial experiente ache que precise melhorar e ajude-o com ideias e estímulos para reflexões.`,
			},
	}

	messages = append(messages, oldMessagesFormated...)
	messages = append(messages, dto.ChatGptMessage{Role: "user", Content: userMessage,})
 
	jsonPayload, err := json.Marshal(dto.NewChatGptMessageDto(messages))

	if err != nil {
		return models.MessageModel{}, errors.New(err.Error())
	}
 
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))

	if err != nil {
		return models.MessageModel{}, errors.New(err.Error())
	}
 
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+bearerToken)
 
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Printf("Erro ao fazer a request %s", err.Error())
		return models.MessageModel{}, errors.New(err.Error())
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return models.MessageModel{}, errors.New(err.Error())
	}
 
	var response dto.ChatGPTResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return models.MessageModel{}, errors.New(err.Error())
	}

	if len(response.Choices) == 0 {
		return models.MessageModel{}, errors.New("nothing return")
	}	

	var gptMessage models.MessageModel = models.MessageModel {
		Content: response.Choices[len(response.Choices)-1].Message.Content,
		UserID: 1,
		ChatRoomId: chatRoomId,
		CreatedAt: time.Now(),
	}

	return gptMessage, nil
}