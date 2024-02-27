package system

import (
	"app/model"
)

type JwtBlacklist struct {
	model.Model
	Jwt string `gorm:"type:text;comment:jwt"`
}
