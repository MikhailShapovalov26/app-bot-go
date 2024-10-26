package app

import (
	config "alert/configs"
	"alert/pkg/api"

	"fmt"
	"net/http"
)

func CreateServer(cfg *config.Config) {
	fmt.Printf("Server is running on %s:%s\n", cfg.Telegram.Host, cfg.Telegram.Port)

	http.HandleFunc("/", testHandlers)
	http.HandleFunc("/message", api.MessageHandler(cfg))
	err := http.ListenAndServe(cfg.Telegram.Host+":"+cfg.Telegram.Port, nil)
	if err != nil {
		panic(err)
	}

}
func testHandlers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ОК")
}
