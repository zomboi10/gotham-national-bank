package usecase

import (
	"context"

	"github.com/zomboi10/GNB/internal/entity"
)

type AccountRepository interface {
	Create(ctx context.Context, req entity.CreateAccountRequest) (entity.Account, error)
	List(ctx context.Context, req entity.ListAccountsRequest) ([]entity.Account, error)
	Get(ctx context.Context, id int64) (entity.Account, error)
	Update(ctx context.Context, req entity.UpdateAccountRequest) (entity.Account, error)
	Delete(ctx context.Context, id int64) error
	AddBalance(ctx context.Context, req entity.AddAccountBalance) (entity.Account, error)
}
