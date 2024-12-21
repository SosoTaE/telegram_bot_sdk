package main

type User struct {
	Id        int64  `json:"id"`
	IsBot     bool   `json:"is_bot"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}

type Chat struct {
	Id        int64  `json:"id"`
	Type      string `json:"type"`
	Title     string `json:"title"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Message struct {
	MessageId  int64  `json:"message_id"`
	From       User   `json:"from"`
	SenderChat Chat   `json:"sender_chat"`
	Text       string `json:"text"`
}

type MessageId struct {
	Chat      Chat `json:"chat"`
	MessageId int  `json:"message_id"`
	Date      int  `json:"date"`
}

type SendMessageType struct {
	ChatID string `json:"chat_id"`
	Text   string `json:"text"`
}

type InputMediaPhoto struct {
	Type    string `json:"type"`
	Media   string `json:"media"`
	Caption string `json:"caption"`
}

type SendMediaGroupType struct {
	ChatID string            `json:"chat_id"`
	Media  []InputMediaPhoto `json:"media"`
}
