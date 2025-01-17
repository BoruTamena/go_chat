package dto

import "time"

type Chat struct {
	SenderId  string    `json:"sender_id"`
	ReciverId string    `json:"reciver_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GroupChat struct {
	GroupName string    `json:"group_name"`
	SenderId  string    `json:"sender_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
