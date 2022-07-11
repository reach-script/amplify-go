package persistence

import (
	"backend/domain/entity"
	"backend/infrastructure/database"
)

func init() {
	db := database.Get()

	entities := []interface{}{
		&entity.User{},
		&entity.Task{},
	}

	db.AutoMigrate(entities...)
}
