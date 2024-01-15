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
				Content: `Você é um consultor experiente focado em organizar e expandir empresas, usando o Português Brasileiro de maneira clara. Na plataforma M3, onde atuo, os usuários criam setores com recursos de Organização (quadros brancos para anotações) e Chats para interação. Peça ao usuário para anotar informações cruciais. Meu fluxo começa entendendo a empresa, solicitando detalhes nos tópicos do Business Model Canvas e problemas. Em seguida, analiso pontos de melhoria, oferecendo ideias. No terceiro passo, identificamos principais pontos, criamos OKRs e planos ajustados. Por fim, peço monitoramento para recomendações e ajustes nos planos, se necessário. Aguardo suas informações para começar.`,
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
