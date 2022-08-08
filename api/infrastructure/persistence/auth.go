package persistence

import (
	"backend/domain/entity"
	"backend/domain/repository"
	"backend/packages/errors"
	"log"

	"github.com/guregu/dynamo"
)

type authPersistance struct{}

func NewAuthPersistance() repository.AuthRepository {
	return &authPersistance{}
}

func (a *authPersistance) Logout(db *dynamo.DB, auth *entity.Auth) errors.IError {
	table := db.Table("Session")

	log.Println(auth)
	table.Get("key1", auth.Sub).Range("key2", dynamo.Equal, auth.Payload).One(&auth)
	auth.Disabled = true
	if err := table.Put(auth).Run(); err != nil {
		return errors.NewUnexpectedError(err)
	}
	return nil
}
