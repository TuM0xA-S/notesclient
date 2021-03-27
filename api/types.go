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
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	UserID    int       `json:"user_id"`
	Published bool      `json:"published"`
}

// User ...
type User struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
}

// NotePatch with nullable fields
type NotePatch struct {
	Body      string `json:"body"`
	Title     string `json:"title"`
	Published bool   `json:"published"`
	ID        int    `json:"-"`
}
