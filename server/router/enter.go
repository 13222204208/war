package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
	"github.com/flipped-aurora/gin-vue-admin/server/router/war"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	War     war.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
