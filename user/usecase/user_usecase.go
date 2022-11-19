package usecase

import (
	"context"
	"exercise/domain"
	"exercise/domain/entity"
)

type userUsecase struct {
	UserRepository entity.UserRepository
}

func NewUserUsecase(userRepository entity.UserRepository) *userUsecase {
	return &userUsecase{UserRepository: userRepository}
}

func (u *userUsecase) Register(ctx context.Context, userRegister *entity.UserRegister) (string, error) {
	_, err := u.UserRepository.FindByEmail(ctx, userRegister.Email)
	if err == nil {
		return "", domain.ErrEmailAlreadyExist
	}

	user, err := entity.NewUser(userRegister.Email, userRegister.Name, userRegister.Password)
	if err != nil {
		return "", err
	}

	if err = u.UserRepository.Create(ctx, user); err != nil {
		return "", err
	}

	token, err := user.GenerateJWT()
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *userUsecase) Login(ctx context.Context, userLogin *entity.UserLogin) (string, error) {
	user, err := u.UserRepository.FindByEmail(ctx, userLogin.Email)
	if err != nil {
		return "", err
	}

	if correctPassword := user.CorrectPassword(userLogin.Password); !correctPassword {
		return "", domain.ErrUnauthorized
	}

	token, err := user.GenerateJWT()
	if err != nil {
		return "", err
	}

	return token, nil
}
