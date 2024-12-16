package entity

import "time"

type (
	Transfer struct {
		ID            int64     `json:"id"`
		FromAccountID int64     `json:"from_account_id"`
		ToAccountID   int64     `json:"to_account_id"`
		Amount        int64     `json:"amount"` // must be positive
		CreatedAt     time.Time `json:"created_at"`
	}
)
