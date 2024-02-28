package v1

import (
	"github.com/weijianhong/igo/api/v1/base"
	"github.com/weijianhong/igo/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup system.ApiGroup
	BaseApiGroup   base.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
