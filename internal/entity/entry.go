package entity

import "time"

type (
	Entry struct {
		ID        int64     `json:"id"`
		AccountID int64     `json:"account_id"`
		Amount    int64     `json:"amount"` // can be negative or positive
		CreatedAt time.Time `json:"created_at"`
	}
)
