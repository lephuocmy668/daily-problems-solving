package repositories

import (
	"github.com/lephuocmy668/daily-problems-solving/golang/workspace/tiktok-clean-microservice/businesslogics/domains"
)

type UserRequest struct {
}

type UserRepository interface {
	Store(u *domains.User) (*domains.User, error)
	Find(req UserRequest) ([]domains.User, error)
	FindOne(req UserRequest) (*domains.User, error)
	Update(u *domains.User) (*domains.User, error)
}
