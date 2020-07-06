package system

import (
	v1 "gf-server/app/api/v1"
	"gf-server/global"
	"gf-server/middleware"
)

// InitJwtRouter 注册jwt相关路由
func InitJwtRouter() {
	// TODO 缺少CasbinHandler中间件
	ApiRouter := global.GFVA_SERVER.Group("jwt").Middleware(middleware.MiddlewareAuth)
	{
		ApiRouter.POST("jsonInBlacklist", v1.JsonInBlacklist) // jwt加入黑名单
	}
}
