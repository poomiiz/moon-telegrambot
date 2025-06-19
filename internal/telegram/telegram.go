package telegram

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"moon-telegrambot/internal/ai"
	"net/http"
	"os"
)

func RunServer(addr string) {
	http.HandleFunc("/webhook", webhookHandler)
	http.HandleFunc("/healthz", healthHandler)
	log.Println("Listening on", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

type TelegramUpdate struct {
	Message struct {
		Chat struct {
			ID int64 `json:"id"`
		} `json:"chat"`
		Text string `json:"text"`
	} `json:"message"`
}

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	log.Println("TELEGRAM webhook received:", string(body))

	var update TelegramUpdate
	json.Unmarshal(body, &update)

	if update.Message.Text != "" {
		reply, _ := ai.GetAIReply(update.Message.Chat.ID, update.Message.Text)
		sendMessage(update.Message.Chat.ID, reply)
	}
	w.WriteHeader(200)
}

func sendMessage(chatID int64, msg string) {
	body := map[string]interface{}{
		"chat_id": chatID,
		"text":    msg,
	}
	b, _ := json.Marshal(body)
	log.Println("Sending telegram reply:", string(b))
	http.Post("https://api.telegram.org/bot"+os.Getenv("TELEGRAM_BOT_TOKEN")+"/sendMessage", "application/json", bytes.NewReader(b))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}
