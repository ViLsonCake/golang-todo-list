package entity

import "time"

type User struct {
	ID       uint      `json:"id" gorm:"primary_key"`
	Username string    `json:"username" gorm:"unique"`
	Email    string    `json:"email" gorm:"unique"`
	Password string    `json:"-"`
	Created  time.Time `json:"created"`
}

type Todo struct {
	ID      uint      `json:"id" gorm:"primary_key"`
	Text    string    `json:"text"`
	Done    bool      `json:"done"`
	Created time.Time `json:"created"`
	UserId  uint      `json:"user_id"`
}
