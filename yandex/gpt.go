package yandex

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const apiKey = "AQVNxr_b_2Nhqw1nEzSqB0zXUgeJiG68vrgR6LOD"
const apiURL = "https://api.yandexgpt.com/v1/completion"

type YandexGPTRequest struct {
	Model             string                     `json:"model"`
	InstructionURI    string                     `json:"instruction_uri"`
	RequestText       string                     `json:"request_text"`
	GenerationOptions YandexGPTGenerationOptions `json:"generation_options"`
}

type YandexGPTGenerationOptions struct {
	MaxTokens   int     `json:"max_tokens"`
	Temperature float64 `json:"temperature"`
}

type YandexGPTResponse struct {
	Text string `json:"text"`
}

func GetYandexGPTResponse() {
	requestBody := YandexGPTRequest{
		Model:          "general",
		InstructionURI: "ds://bt120djph01n8cr7ct2l",
		RequestText:    "Как приобрести услуги клуба?",
		GenerationOptions: YandexGPTGenerationOptions{
			MaxTokens:   1000,
			Temperature: 0.1,
		},
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatalf("Error marshaling request: %v", err)
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		log.Fatalf("Error response from API: %s", string(body))
	}

	var response YandexGPTResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Fatalf("Error decoding response: %v", err)
	}

	fmt.Println("Response:", response.Text)
}
