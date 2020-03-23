package users

import (
	"github.com/jinzhu/gorm"
)

type UserInfo struct {
	gorm.Model
	Username           string
	Password           string
	SuperUser          string
	AccountPerComputer string
}
