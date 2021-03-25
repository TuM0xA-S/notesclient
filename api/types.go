package api

import (
	"time"
)

// ResponseData ...
type ResponseData struct {
	Success     bool           `json:"success"`
	Message     string         `json:"message"`
	Notes       []Note         `json:"notes"`
	AccessToken string         `json:"access_token"`
	Note        Note           `json:"note"`
	User        User           `json:"user"`
	Pagination  map[string]int `json:"pagination"`
}

// Note ...
type Note struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	UserID    uint      `json:"user_id"`
	Published bool      `json:"published"`
}

// User ...
type User struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
}
