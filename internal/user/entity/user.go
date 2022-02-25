package entity

import (
	"time"
)

type User struct {
	UserId    string `json:"id""`
	AvatarURL string `json:"avatarURL" `
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	// Some other DB related attributes
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
