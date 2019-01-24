package usecases

import (
	"github.com/lephuocmy668/daily-problems-solving/golang/workspace/tiktok-clean-microservice/businesslogics/repositories"
)

// UseCases ...
type UseCases struct {
	User UserInteractor
}

func MakeUseCases(userRepository repositories.UserRepository) *UseCases {
	return &UseCases{
		User: UserInteractor{
			userRepository,
		},
	}
}
