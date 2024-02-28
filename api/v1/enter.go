package v1

import (
	"github.com/weijianhong/igo/api/v1/system"
)

type ApiGroup struct {
	system.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
