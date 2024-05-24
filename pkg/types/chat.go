package types

import (
	"context"
	"fmt"

	"github.com/sheeiavellie/go-yandexgpt"
)

func main() {
	client := yandexgpt.NewYandexGPTClientWithAPIKey("apiKey")
	request := yandexgpt.YandexGPTRequest{
		ModelURI: yandexgpt.MakeModelURI("catalogID", yandexgpt.YandexGPTModelLite),
		CompletionOptions: yandexgpt.YandexGPTCompletionOptions{
			Stream:      false,
			Temperature: 0.7,
			MaxTokens:   2000,
		},
		Messages: []yandexgpt.YandexGPTMessage{
			{
				Role: yandexgpt.YandexGPTMessageRoleSystem,
				Text: "Every time you get ONE you answer just TWO",
			},
			{
				Role: yandexgpt.YandexGPTMessageRoleUser,
				Text: "ONE",
			},
		},
	}

	response, err := client.CreateRequest(context.Background(), request)
	if err != nil {
		fmt.Println("Error")
		return
	}

	fmt.Println(response.Result.Alternatives[0].Message.Text)
}
