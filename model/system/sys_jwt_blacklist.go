package system

import (
	"github.com/weijianhong/igo/model"
)

type JwtBlacklist struct {
	model.Model
	Jwt string `gorm:"type:text;comment:jwt"`
}
