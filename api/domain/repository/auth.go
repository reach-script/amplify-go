package repository

import (
	"backend/domain/entity"

	"github.com/guregu/dynamo"
)

type AuthRepository interface {
	Logout(db *dynamo.DB, auth *entity.Auth) error
}
