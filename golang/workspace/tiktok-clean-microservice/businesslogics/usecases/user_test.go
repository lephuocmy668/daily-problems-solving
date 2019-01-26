package usecases

import (
	"reflect"
	"testing"

	"github.com/lephuocmy668/daily-problems-solving/golang/workspace/tiktok-clean-microservice/businesslogics/domains"
	"github.com/lephuocmy668/daily-problems-solving/golang/workspace/tiktok-clean-microservice/businesslogics/repositories"
)

type UserIndex struct {
}

func (userIndex *UserIndex) Store(u *domains.User) (*domains.User, error) {
	return nil, nil
}

func (userIndex *UserIndex) Find(req repositories.UserRequest) ([]domains.User, error) {
	return nil, nil
}

func (userIndex *UserIndex) FindOne(req repositories.UserRequest) (*domains.User, error) {
	return nil, nil
}

func (userIndex *UserIndex) Update(u *domains.User) (*domains.User, error) {
	return nil, nil
}

func TestUserInteractor_Create(t *testing.T) {
	type fields struct {
		UserRepository repositories.UserRepository
	}
	type args struct {
		u *domains.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domains.User
		wantErr bool
	}{
		{
			name: "Create should return new user info",
			fields: fields{
				UserRepository: &UserIndex{},
			},
			args: args{
				u: &domains.User{},
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ui := &UserInteractor{
				UserRepository: tt.fields.UserRepository,
			}
			got, err := ui.Create(tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserInteractor.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserInteractor.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
