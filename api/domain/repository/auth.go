package repository

import (
	"backend/domain/entity"
	"backend/packages/errors"

	"github.com/guregu/dynamo"
)

type AuthRepository interface {
	Logout(db *dynamo.DB, auth *entity.Auth) errors.IError
}
