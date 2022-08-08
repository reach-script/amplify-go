package usecase

import (
	"backend/domain/entity"
	"backend/domain/repository"
	"backend/packages/errors"

	"github.com/guregu/dynamo"
)

type Auth interface {
	Logout(db *dynamo.DB, auth *entity.Auth) errors.IError
}

type authUseCase struct {
	authRepository repository.AuthRepository
}

func NewAuthUseCase(authRepository repository.AuthRepository) Auth {
	return &authUseCase{
		authRepository: authRepository,
	}
}

func (a *authUseCase) Logout(db *dynamo.DB, auth *entity.Auth) errors.IError {
	if err := a.authRepository.Logout(db, auth); err != nil {
		return errors.NewUnexpectedError(err)
	}
	return nil
}
