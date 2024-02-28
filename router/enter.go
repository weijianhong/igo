package router

import (
	"github.com/gin-gonic/gin"
	"github.com/weijianhong/igo/router/system"
)

type RouterGroup struct {
	pub, pri, ws *gin.RouterGroup
	System       system.RouterGroup
}

var RouterGroupApp = new(RouterGroup)

func (a RouterGroup) Add(pub, pri, ws *gin.RouterGroup) {
	a.pub = pub
	a.pri = pri
	a.ws = ws

	a.System.InitSystemRouter(a.pri)

	return
}
