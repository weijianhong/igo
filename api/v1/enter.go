package v1

import (
	"app/api/v1/system"
)

type ApiGroup struct {
	system.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
