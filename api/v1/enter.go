package v1

import (
	"github.com/weijianhong/igo/api/v1/system"
)

type ApiGroup struct {
	system.ApiSystemGroup
}

var ApiGroupApp = new(ApiGroup)
