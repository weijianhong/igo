package router

import (
	"github.com/gin-gonic/gin"
	"github.com/weijianhong/igo/router/system"
)

type group struct{}

var GroupNew = new(group)

func (a group) Add(pub, pri, ws *gin.RouterGroup) {

	system.GroupNew.Add(pri)

	return
}
