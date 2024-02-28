package v1

import (
	"github.com/weijianhong/igo/api/v1/base"
	"github.com/weijianhong/igo/api/v1/system"
)

type ApiGroup struct {
	system.ApiSystemGroup
	base.ApiBaseGroup
}

var ApiGroupApp = new(ApiGroup)
