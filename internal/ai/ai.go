package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func GetAIReply(chatID interface{}, msg string) (string, error) {
	payload := map[string]interface{}{
		"userId": fmt.Sprintf("%v", chatID),
		"text":   msg,
	}
	b, _ := json.Marshal(payload)
	log.Println("Calling AI-service with:", string(b))
	req, _ := http.NewRequest("POST", os.Getenv("AI_SERVICE_URL")+"/chat", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error calling AI-service:", err)
		return "ระบบขัดข้อง ลองใหม่อีกครั้ง", err
	}
	defer resp.Body.Close()
	var out struct {
		Reply string `json:"reply"`
	}
	json.NewDecoder(resp.Body).Decode(&out)
	log.Println("AI-service replied:", out.Reply)
	return out.Reply, nil
}
