package telegram_bot_sdk

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type Bot struct {
	Token string
}

func (bot *Bot) SendMessage(message SendMessageType) (map[string]interface{}, error) {
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
