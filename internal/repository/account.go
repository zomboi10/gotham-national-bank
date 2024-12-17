package repository

import (
	"context"
	"database/sql"

	"github.com/zomboi10/GNB/internal/entity"
)

type AccountRepository struct {
	*sql.DB
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{db}
}

func (r AccountRepository) AddBalance(ctx context.Context, req entity.AddAccountBalance) (entity.Account, error) {
	query := `UPDATE accounts SET balance = balance + $1 WHERE id = $2 RETURNING id, owner, balance, currency, created_at`
	row := r.QueryRowContext(ctx, query, req.Amount, req.ID)

	var account entity.Account
	err := row.Scan(
		&account.ID,
		&account.Owner,
		&account.Balance,
		&account.Currency,
		&account.CreatedAt,
	)
	return account, err
}

func (r AccountRepository) Create(ctx context.Context, req entity.CreateAccountRequest) (entity.Account, error) {
	var account entity.Account

	query := `INSERT INTO accounts (owner, balance, currency) VALUES ($1, $2, $3) RETURNING id, owner, balance, currency, created_at`
	row := r.QueryRowContext(ctx, query, req.Owner, req.Balance, req.Currency)

	err := row.Scan(
		&account.ID,
		&account.Owner,
		&account.Balance,
		&account.Currency,
		&account.CreatedAt,
	)
	return account, err
}

func (r AccountRepository) List(ctx context.Context, req entity.ListAccountsRequest) ([]entity.Account, error) {
	query := `SELECT id, owner, balance, currency, created_at FROM accounts WHERE owner = $1 ORDER BY id LIMIT $2 OFFSET $3`
	rows, err := r.QueryContext(ctx, query, req.Owner, req.Limit, req.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var accounts []entity.Account
	for rows.Next() {
		var account entity.Account
		if err := rows.Scan(
			&account.ID,
			&account.Owner,
			&account.Balance,
			&account.Currency,
			&account.CreatedAt,
		); err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return accounts, nil
}

func (r AccountRepository) Get(ctx context.Context, id int64) (entity.Account, error) {
	query := `SELECT id, owner, balance, currency, created_at FROM accounts WHERE id = $1 LIMIT 1 FOR NO KEY UPDATE`
	row := r.QueryRowContext(ctx, query, id)

	var account entity.Account
	err := row.Scan(
		&account.ID,
		&account.Owner,
		&account.Balance,
		&account.Currency,
		&account.CreatedAt,
	)
	return account, err
}

func (r AccountRepository) Update(ctx context.Context, req entity.UpdateAccountRequest) (entity.Account, error) {
	query := `UPDATE acccounts SET balance = $2 WHERE id = $1 RETURNING id, owner, balance, currency, created_at`
	row := r.QueryRowContext(ctx, query, req.ID, req.Balance)

	var account entity.Account
	err := row.Scan(
		&account.ID,
		&account.Owner,
		&account.Balance,
		&account.Currency,
		&account.CreatedAt,
	)
	return account, err
}

func (r AccountRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM accounts WHERE id = $1`
	_, err := r.ExecContext(ctx, query, id)
	return err
}
