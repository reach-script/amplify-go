package persistence

import (
	"backend/domain/entity"
	"backend/domain/repository"
	"log"

	"github.com/guregu/dynamo"
)

type authPersistance struct{}

func NewAuthPersistance() repository.AuthRepository {
	return &authPersistance{}
}

func (a *authPersistance) Logout(db *dynamo.DB, auth *entity.Auth) error {
	table := db.Table("Session")

	log.Println(auth)
	table.Get("key1", auth.Key1).Range("key2", dynamo.Equal, auth.Key2).One(&auth)
	auth.Disabled = true
	if err := table.Put(auth).Run(); err != nil {
		return err
	}
	return nil
}
