package entity

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name       string `json:"name" gorm:"column:name;"`
	CognitoSub string `json:"cognito_sub" gorm:"column:cognito_sub;"`
	Tasks      []Task `json:"tasks" gorm:"foreignKey:UserId;references:UserId"`
}
