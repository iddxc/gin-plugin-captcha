package model

import "time"

type Message struct {
	TS      time.Time `json:"ts"`
	Sender  string    `json:"sender"`
	Message string    `json:"message"`
}

type MessageReq struct {
	Sender     string   `json:"sender"`
	Message    string   `json:"message"`
	ChatId     string   `json:"chat_id"`
	RoomName   string   `json:"room_name"`
	Recipients []string `json:"recipients"`
}
