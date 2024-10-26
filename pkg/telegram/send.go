package telegram

import (
	config "alert/configs"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func SendToTelegram(cfg *config.Config, message, author string) error {

	client := &http.Client{}

	payload := map[string]string{
		"chat_id": cfg.Telegram.Chat_id,
		"text":    fmt.Sprintf("Message from %s: %s", author, message),
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %v", err)
	}
	fmt.Println(string(jsonPayload))
	req, err := http.NewRequest("POST", "https://api.telegram.org/bot"+cfg.Telegram.Token+"/sendMessage", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return fmt.Errorf("failed to send message to Telegram: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	// Выполняем запрос
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}
	fmt.Printf("%s\n", bodyText)
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Unexpected status code: %d, body: %s", resp.StatusCode, bodyText)
	}
	return nil
}
