package usecases

import (
	"github.com/lephuocmy668/daily-problems-solving/golang/workspace/tiktok-clean-microservice/businesslogics/domains"
	"github.com/lephuocmy668/daily-problems-solving/golang/workspace/tiktok-clean-microservice/businesslogics/repositories"
)

type UserInteractor struct {
	UserRepository repositories.UserRepository
}

func (ui *UserInteractor) Create(u *domains.User) (*domains.User, error) {
	return ui.UserRepository.Store(u)
}

func (ui *UserInteractor) Update(u *domains.User) (*domains.User, error) {
	return ui.UserRepository.Update(u)
}

func (ui *UserInteractor) FindOne(req repositories.UserRequest) (*domains.User, error) {
	return ui.UserRepository.FindOne(req)
}

func (ui *UserInteractor) Find(req repositories.UserRequest) ([]domains.User, error) {
	return ui.UserRepository.Find(req)
}
