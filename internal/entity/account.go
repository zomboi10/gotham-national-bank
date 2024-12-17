package entity

import "time"

type (
	Account struct {
		ID        int64     `json:"id"`
		Owner     string    `json:"owner"`
		Balance   int64     `json:"balance"`
		Currency  string    `json:"currency"`
		CreatedAt time.Time `json:"created_at"`
	}

	CreateAccountRequest struct {
		Owner    string `json:"owner"`
		Balance  int64  `json:"balance"`
		Currency string `json:"currency"`
	}

	ListAccountsRequest struct {
		Owner  string `json:"owner"`
		Limit  int32  `json:"limit"`
		Offset int32  `json:"offset"`
	}

	UpdateAccountRequest struct {
		ID      int64 `json:"id"`
		Balance int64 `json:"balance"`
	}

	AddAccountBalance struct {
		Amount int64 `json:"amount"`
		ID     int64 `json:"id"`
	}
)
