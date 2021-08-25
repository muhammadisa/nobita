package _interface

import (
	"context"
	"github.com/muhammadisa/nobita/model/v1/auth"
)

type RW interface {
	WriteAccount(context.Context, auth.Account) (int64, error)
	WriteProfile(context.Context, int64) error
	ReadIsAccountExist(context.Context, auth.Account) (int64, error)
	ReadProfileByAccountID(ctx context.Context, accountID int64) (auth.Profile, error)
	ReadContactByID(context.Context, int64) (auth.Contact, error)
	UpdateProfileByAccountID(context.Context, auth.Profile)error
	UpdateRoleAndDeviceID(context.Context, int64, int64, int64) error
	UpdateLoginStatus(context.Context, int64, bool) error
	VerifyTempSecret(context.Context, auth.Secret) (auth.Account, error)
	UpdateTempSecret(context.Context, int64, string) error
}
