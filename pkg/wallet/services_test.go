package wallet

import (
	"testing"

	"github.com/Sher-mad/wallet/pkg/types"
	"github.com/google/uuid"
)

var (
	phone_1 = types.Phone("+992919558857")
	phone_2 = types.Phone("+992988099969")
	phone_3 = types.Phone("+992551000991")
	phone_4 = types.Phone("+992208099969")
)

func TestService_FindAccountByID_success(t *testing.T) {
	svc := &Service{}
	_, err := svc.RegisterAccount(phone_1)
	if err != nil {
		t.Error(err)
		return
	}
	_, err = svc.RegisterAccount(phone_2)
	if err != nil {
		t.Error(err)
		return
	}
	_, err = svc.RegisterAccount(phone_3)
	if err != nil {
		t.Error(err)
		return
	}
	_, err = svc.RegisterAccount(phone_4)
	if err != nil {
		t.Error(err)
		return
	}
	account, err := svc.FindAccountByID(1)
	if err != nil {
		t.Error(err, " = ", account)
		return
	}
	err = svc.Deposit(1, 10_000_00)
	if err != nil {
		t.Error(err)
		return
	}
	account, err = svc.FindAccountByID(2)
	if err != nil {
		t.Error(err, " = ", account)
		return
	}
	err = svc.Deposit(2, 20_000_00)
	if err != nil {
		t.Error(err)
		return
	}
	account, err = svc.FindAccountByID(3)
	if err != nil {
		t.Error(err, " = ", account)
		return
	}
	err = svc.Deposit(3, 30_000_00)
	if err != nil {
		t.Error(err)
		return
	}
	account, err = svc.FindAccountByID(4)
	if err != nil {
		t.Error(err, " = ", account)
		return
	}
	err = svc.Deposit(4, 10_000_00)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestService_FindAccountByID_fail(t *testing.T) {
	svc := &Service{}
	_, err := svc.RegisterAccount(phone_1)
	if err != nil {
		t.Error(err)
		return
	}
	_, err = svc.RegisterAccount(phone_2)
	if err != nil {
		t.Error(err)
		return
	}
	_, err = svc.RegisterAccount(phone_3)
	if err != nil {
		t.Error(err)
		return
	}
	_, err = svc.RegisterAccount(phone_4)
	if err != nil {
		t.Error(err)
		return
	}
	account, err := svc.FindAccountByID(int64(uuid.New().ID()))
	if err == nil {
		t.Error(err, " = ", account)
		return
	}
	err = svc.Deposit(int64(uuid.New().ID()), 10_000_00)
	if err == nil {
		t.Error(err)
		return
	}
}
