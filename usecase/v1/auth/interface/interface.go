package _interface

import (
	"context"
	"github.com/muhammadisa/nobita/model/v1/auth"
)

type UseCase interface {
	AuthAccount(context.Context, auth.Account) (auth.Status, error)
	VerifySecret(context.Context, auth.Secret) (auth.Verificator, error)
	MyProfile(context.Context, int64) (auth.Profile, error)
	EditProfile(context.Context, auth.Profile) error
}
