package auth

import (
	"context"
	"github.com/muhammadisa/nobita/model/v1/auth"
)

func (u usecase) EditProfile(ctx context.Context, profile auth.Profile) error {
	return u.AuthRepository.UpdateProfileByAccountID(ctx, profile)
}
