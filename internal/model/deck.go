package model

import "time"

type Deck struct {
	ID        int    `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	ChatID    int64  `json:"chatId" db:"chatId"`
	CreatedAt time.Time   `json:"createdAt" db:"createdAt"`
}