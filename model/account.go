package model

import "time"

type Account struct {
	ID         uint      `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	AppVersion string    `json:"app_version"`
	Email      string    `json:"email"`
	IsDeleted  bool      `json:"is_deleted"`
}
