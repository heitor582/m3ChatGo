package dto

type ChatGptMessageDto struct {
	Model	string `json:"model" validate:"required,string"`
	Temperature float32 `json:"temperature" validate:"required"`
	Messages []ChatGptMessage `json:"messages" validate:"required"`
}

type ResponseFormat struct{
	Type string `json:"type" validate:"required,string"`
}

type ChatGptMessage struct {
	Role	string `json:"role" validate:"required,string"`
	Content string `json:"content" validate:"required,string"`
}

func NewChatGptMessageDto(newMessage ChatGptMessage) ChatGptMessageDto {
	return ChatGptMessageDto{
		Model: "gpt-3.5-turbo",
		Temperature: 1.0,
		Messages: []ChatGptMessage{
			{
				Role: "system",
				Content: `Você é um consultor empresarial experiente, que trabalha com foco em organizar e crescer empresas.
				Você usa o idioma Português Brasileiro e tem uma comunicação clara e objetiva. Além de estar em uma posição que guia a conversa para a organização e crescimento constante da empresas.
				Seu fluxo de trabalho é primeiramente entender a empresa do cliente ao completo e depois orienta-lo. 
				Primeiro peça informações sobre sua empresa como os tópicos do Business Model Canva. (Detalhe o que é cada tópicos)
				Segundo analise os pontos que um consultor empresarial experiente ache que precise melhorar e ajude-o com ideias e estímulos para reflexões.`,
			},
			newMessage,
		},
	}
}