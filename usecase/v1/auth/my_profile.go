package auth

import (
	"context"
	"github.com/muhammadisa/nobita/model/v1/auth"
)

func (u usecase) MyProfile(ctx context.Context, accountID int64) (auth.Profile, error) {
	return u.AuthRepository.ReadProfileByAccountID(ctx, accountID)
}
