package searchengine

import (
	"github.com/lephuocmy668/daily-problems-solving/golang/workspace/tiktok-clean-microservice/businesslogics/domains"
	"github.com/lephuocmy668/daily-problems-solving/golang/workspace/tiktok-clean-microservice/businesslogics/repositories"
)

type UserIndex struct {
}

func (userIndex *UserIndex) Store(u *domains.User) (*domains.User, error) {
	return u, nil
}

func (userIndex *UserIndex) Find(req repositories.UserRequest) ([]domains.User, error) {
	return nil, nil
}

func (userIndex *UserIndex) FindOne(req repositories.UserRequest) (*domains.User, error) {
	user := domains.User{
		FirstName: "My",
		LastName:  "Le Phuoc",
	}

	return &user, nil
}

func (userIndex *UserIndex) Update(u *domains.User) (*domains.User, error) {
	return nil, nil
}
