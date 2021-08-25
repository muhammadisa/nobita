package auth

import (
	authrwinterface "github.com/muhammadisa/nobita/repository/v1/auth/interface"
	authusecaseinterface "github.com/muhammadisa/nobita/usecase/v1/auth/interface"
	"github.com/muhammadisa/nobita/util/otp"
)

type usecase struct {
	AuthRepository authrwinterface.RW
	Sender         otp.OTP
}

func NewAuthUseCaseV1(authRepository authrwinterface.RW, config otp.Config) authusecaseinterface.UseCase {
	return &usecase{
		AuthRepository: authRepository,
		Sender:         otp.NewOTP(config),
	}
}
