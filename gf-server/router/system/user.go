package system

import (
	v1 "gf-server/app/api/v1"
	"gf-server/global"
	"gf-server/middleware"
)

// InitUserRouter 注册用户路由
func InitUserRouter() {
	// TODO 缺少CasbinHandler中间件
	UserRouter := global.GFVA_SERVER.Group("user").Middleware(middleware.MiddlewareAuth)
	{
		UserRouter.POST("changePassword", v1.ChangePassword)   // 修改密码
		UserRouter.POST("uploadHeaderImg", v1.UploadHeaderImg)  // 上传头像
		UserRouter.POST("getUserList", v1.GetUserList)      // 分页获取用户列表
		UserRouter.POST("setUserAuthority", v1.SetUserAuthority) // 设置用户权限
		UserRouter.DELETE("deleteUser", v1.DeleteUser)     // 删除用户
	}
}
