package persistence

import (
	"backend/domain/entity"
	"backend/domain/repository"

	"github.com/guregu/dynamo"
)

type authPersistance struct{}

func NewAuthPersistance() repository.AuthRepository {
	return &authPersistance{}
}

func (a *authPersistance) Logout(db *dynamo.DB, auth *entity.Auth) error {
	table := db.Table("Token")

	table.Get("payload", auth.Payload).One(&auth)
	auth.Disabled = true
	if err := db.Table("Token").Put(auth).Run(); err != nil {
		return err
	}
	return nil
}
