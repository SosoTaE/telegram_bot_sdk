package telegram_bot_sdk

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

type Bot struct {
	Token string
}

func (bot *Bot) sendMessage(message SendMessageType) (map[string]interface{}, error) {
	requestBody, err := json.Marshal(message)
	if err != nil {
		return map[string]interface{}{}, err
	}

	resp, err := http.Post("https://api.telegram.org/bot"+bot.Token+"/sendMessage", "application/json", bytes.NewBuffer(requestBody))

	if err != nil {
		return map[string]interface{}{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return map[string]interface{}{}, err
	}

	var object map[string]interface{}
	err = json.Unmarshal(body, &object)

	if err != nil {
		return map[string]interface{}{}, err
	}

	return object, nil
}

func (bot *Bot) SendMessage(message SendMessageType) (map[string]interface{}, error) {
	if len(message.Text) > 4096 {
		text := message.Text[:4096]
		lastIndex := strings.LastIndex(text, " ")
		text = text[:lastIndex]

		newMessage := SendMessageType{
			ChatID: message.ChatID,
			Text:   text,
		}

		result, err := bot.sendMessage(newMessage)
		if err != nil {
			return result, err
		}

		return bot.SendMessage(SendMessageType{ChatID: message.ChatID, Text: message.Text[lastIndex+1:]})
	}

	return bot.sendMessage(message)

}

func (bot *Bot) SendMediaGroup(message SendMediaGroupType) (map[string]interface{}, error) {
	requestBody, err := json.Marshal(message)
	if err != nil {
		return map[string]interface{}{}, err
	}

	resp, err := http.Post("https://api.telegram.org/bot"+bot.Token+"/sendMediaGroup", "application/json", bytes.NewBuffer(requestBody))

	if err != nil {
		return map[string]interface{}{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return map[string]interface{}{}, err
	}

	var object map[string]interface{}
	err = json.Unmarshal(body, &object)

	if err != nil {
		return map[string]interface{}{}, err
	}

	return object, nil
}
