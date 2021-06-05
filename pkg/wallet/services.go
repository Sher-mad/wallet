package wallet

import (
	"errors"

	"github.com/Sher-mad/wallet/pkg/types"
)

type Service struct {
	nextAccountID int64
	accounts      []*types.Account
	payments      []*types.Payment
}

var (
	ErrPhoneRegistered      = errors.New("phone already registered")
	ErrAmountMustBePositive = errors.New("amount must be greater than zero")
	ErrAccountNotFound      = errors.New("account not found")
)

func (s *Service) RegisterAccount(phone types.Phone) (*types.Account, error) {
	for _, account := range s.accounts {
		if account.Phone == phone {
			return nil, ErrPhoneRegistered
		}
	}
	s.nextAccountID++
	account := &types.Account{
		ID:      s.nextAccountID,
		Phone:   phone,
		Balance: 0,
	}
	s.accounts = append(s.accounts, account)
	return account, nil
}

func (s *Service) FindAccountByID(accountID int64) (*types.Account, error) {
	var account *types.Account

	for _, accounts := range s.accounts {
		if accounts.ID == accountID {
			account = accounts
			break
		}
	}

	if account == nil {
		return nil, ErrAccountNotFound
	}
	return account, nil
}

func (s *Service) Deposit(accountID int64, amount types.Money) error {
	if amount <= 0 {
		return ErrAmountMustBePositive
	}
	account, err := s.FindAccountByID(accountID)

	if err != nil {
		return ErrAccountNotFound
	}

	account.Balance += amount
	return nil
}
