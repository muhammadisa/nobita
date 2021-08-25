package auth

import (
	"context"
	"fmt"
	"github.com/muhammadisa/nobita/model/v1/auth"
	"github.com/muhammadisa/nobita/util/gen"
	"github.com/muhammadisa/nobita/util/hash"
)

func (u usecase) AuthAccount(ctx context.Context, account auth.Account) (status auth.Status, err error) {
	otp, _ := gen.GRS(6)
	encryptedOtp, err := hash.Hashed(otp)
	if err != nil {
		return status, err
	}

	fmt.Println(otp)

	var target, kind string
	if account.Email != "" {
		kind = "email"
		target = account.Email
	} else {
		kind = "phone"
		target = account.Phone
	}
	fmt.Println(kind)

	var id int64
	accountID, err := u.AuthRepository.ReadIsAccountExist(ctx, account)
	if err != nil {
		return status, err
	}

	if accountID != 0 {
		id = accountID
		err := u.AuthRepository.UpdateRoleAndDeviceID(ctx, id, account.RoleID, account.RoleID)
		if err != nil {
			return status, err
		}
	} else {
		newAccountID, err := u.AuthRepository.WriteAccount(ctx, account)
		if err != nil {
			return status, err
		}
		id = newAccountID
		err = u.AuthRepository.WriteProfile(ctx, id)
		if err != nil {
			return status, err
		}
	}

	err = u.AuthRepository.UpdateTempSecret(ctx, id, string(encryptedOtp))
	if err != nil {
		return status, err
	}

	if kind == "email" {
		err := u.Sender.Email(target, otp)
		if err != nil {
			return status, err
		}
	} else {
		err := u.Sender.SMS(target, otp)
		if err != nil {
			return status, err
		}
	}

	status.Message = fmt.Sprintf("we have sent verification code to %s", target)
	status.Sent = true
	return status, nil
}
