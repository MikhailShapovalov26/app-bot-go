package api

import (
	config "alert/configs"
	"alert/pkg/telegram"
	"encoding/json"
	"fmt"
	"net/http"
)

type UpdateMessageInput struct {
	Message string `json:"message"`
	Author  string `json:"author"`
}

func MessageHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input UpdateMessageInput

		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = telegram.SendToTelegram(cfg, input.Message, input.Author)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "Message received and sent to Telegram")
	}
}
