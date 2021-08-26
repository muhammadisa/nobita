package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/muhammadisa/nobita/model/v1/auth"
	"github.com/muhammadisa/nobita/util/gen"
	"github.com/muhammadisa/nobita/util/hash"
	"github.com/muhammadisa/nobita/util/jwt"
	"time"
)

func (u usecase) VerifySecret(ctx context.Context, secret auth.Secret) (auth.Verificator, error) {
	account, err := u.AuthRepository.VerifyTempSecret(ctx, secret)
	if err != nil {
		return auth.Verificator{}, err
	}

	if account.ID == 0 {
		return auth.Verificator{}, errors.New("account not registered")
	}

	sec, _ := gen.GRS(100)
	encryptedSec, err := hash.Hashed(sec)
	if err != nil {
		return auth.Verificator{}, err
	}
	err = u.AuthRepository.UpdateTempSecret(ctx, account.ID, string(encryptedSec))
	if err != nil {
		return auth.Verificator{}, err
	}

	err = u.AuthRepository.UpdateLoginStatus(ctx, account.ID, true)
	if err != nil {
		return auth.Verificator{}, err
	}

	j := jwt.NewJWTGenerateMode(60*time.Minute, "SECRET")
	data, err := j.Generate(map[string]string{
		"user_id":  fmt.Sprintf("%d", account.ID),
		"role_id":  fmt.Sprintf("%d", account.RoleID),
		"timezone": "Asia/Jakarta",
		"time":     time.Now().String(),
	})
	if err != nil {
		return auth.Verificator{}, err
	}

	profile, err := u.AuthRepository.ReadProfileByAccountID(ctx, account.ID)
	if err != nil {
		return auth.Verificator{}, err
	}

	return auth.Verificator{
		Message:   "code successfully verified",
		Accepted:  true,
		AccountID: account.ID,
		Token:     data.Token,
		ProfileCompleted: profile.FullName != "none" || profile.Gender != "none" ||
			profile.LongLat != "0:0" || profile.BirthDay != "00-00-0000",
	}, nil
}
