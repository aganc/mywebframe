package router

import (
	"airport/controllers/http/admin"
	middleware "airport/middleware/http"
	"github.com/gin-gonic/gin"
)

func HTTP(engine *gin.Engine) {
	engine.Use(middleware.Recover)

	router := engine.Group("/api") // api分组

	// 登陆、登出
	router.POST("/login", admin.Login)
	router.POST("/logout", admin.Logout)
	// 通用中间件
	router.Use(middleware.Cors())
	//router.Use(middleware.JWT)
	routerUser := router.Group("/user") // 用户相关路由组
	{
		routerUser.POST("/list", admin.GetUserList)          // 用户列表查询
		routerUser.POST("/add", admin.CreateUser)            // 添加用户
		routerUser.POST("/update_pwd", admin.UpdateUserInfo) // 用户修改密码
		routerUser.POST("/delete", admin.DeleteUserInfo)     // 用户删除
	}
}
