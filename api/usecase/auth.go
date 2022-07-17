package usecase

import (
	"backend/domain/entity"
	"backend/domain/repository"

	"github.com/guregu/dynamo"
)

type Auth interface {
	Logout(db *dynamo.DB, auth *entity.Auth) error
}

type authUseCase struct {
	authRepository repository.AuthRepository
}

func NewAuthUseCase(authRepository repository.AuthRepository) Auth {
	return &authUseCase{
		authRepository: authRepository,
	}
}

func (a *authUseCase) Logout(db *dynamo.DB, auth *entity.Auth) error {
	if err := a.authRepository.Logout(db, auth); err != nil {
		panic(err)
	}
	return nil
}
