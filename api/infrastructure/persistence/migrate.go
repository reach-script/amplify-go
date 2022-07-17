package persistence

import (
	"backend/domain/entity"
	"backend/infrastructure/database"
)

func init() {
	db := database.GetRDB()

	entities := []interface{}{
		&entity.User{},
		&entity.Task{},
	}

	db.AutoMigrate(entities...)
}
