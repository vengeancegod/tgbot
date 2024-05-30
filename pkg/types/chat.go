package types

import "github.com/sheeiavellie/go-yandexgpt"

type Chat struct {
	Messages []yandexgpt.Response `json:"messages"`
}

func NewChat() *Chat {
	return &Chat{}
}

func (c *Chat) Add(response yandexgpt.Response) {
	c.Messages = append(c.Messages, response)
}

func (c *Chat) AddUserMessage(content string) {
	alternative := yandexgpt.YandexGPTAlternative{
		Message: yandexgpt.YandexGPTMessage{
			Text: content,
		},
		Status: "user",
	}
	response := yandexgpt.YandexGPTResponse{
		Result: yandexgpt.YandexGPTResult{
			Alternatives: []yandexgpt.YandexGPTAlternative{alternative},
		},
	}
	c.Add(&response) // Используем указатель на объект response
}

func (c *Chat) AddAssistantMessage(content string) {
	alternative := yandexgpt.YandexGPTAlternative{
		Message: yandexgpt.YandexGPTMessage{
			Text: content,
		},
		Status: "assistant",
	}
	response := yandexgpt.YandexGPTResponse{
		Result: yandexgpt.YandexGPTResult{
			Alternatives: []yandexgpt.YandexGPTAlternative{alternative},
		},
	}
	c.Add(&response) // Используем указатель на объект response
}

func (c *Chat) AddSystemMessage(content string) {
	alternative := yandexgpt.YandexGPTAlternative{
		Message: yandexgpt.YandexGPTMessage{
			Text: content,
		},
		Status: "system",
	}
	response := yandexgpt.YandexGPTResponse{
		Result: yandexgpt.YandexGPTResult{
			Alternatives: []yandexgpt.YandexGPTAlternative{alternative},
		},
	}
	c.Add(&response) // Используем указатель на объект response
}

func (c *Chat) LastContent() string {
	if len(c.Messages) == 0 {
		return ""
	}
	lastMessage := c.Messages[len(c.Messages)-1]
	yandexResponse, ok := lastMessage.(*yandexgpt.YandexGPTResponse)
	if !ok {
		return ""
	}
	if len(yandexResponse.Result.Alternatives) == 0 {
		return ""
	}
	return yandexResponse.Result.Alternatives[0].Message.Text
}
