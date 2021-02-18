package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.io/iv587/goredis-admin/auth"
	"github.io/iv587/goredis-admin/db"
	"github.io/iv587/goredis-admin/http"
)

var errHandler = func(ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			err1 := err.(error)
			errResJson(ctx, err1.Error(), err1)
		}
	}()
	ctx.Next()
}

var authHander = func(ctx *gin.Context) {
	tokenStr := ctx.PostForm("token")
	ok, user, err := auth.Verify(tokenStr)
	if err != nil {
		panic(err)
	}
	if !ok {
		noLoginJson(ctx)
		ctx.Abort()
	} else {
		u, err := db.GetUserInfo(user.Id)
		if err != nil {
			noLoginWithMsg(ctx, err.Error())
			ctx.Abort()
		}
		if u.Name != user.Name || u.Password != user.Password {
			noLoginWithMsg(ctx, "用户状态已变更，请重新登录！")
			ctx.Abort()
		}
		ctx.Set("userId", user.Id)
		ctx.Next()
	}
}

var RouteMap http.Route = func(en *gin.Engine) {
	en.GET("/", func(c *gin.Context) {
		fmt.Println("hello")
		c.File("web/index.html")
	})
	en.GET("/css/:file", func(c *gin.Context) {
		c.File("web/css/" + c.Params.ByName("file"))
	})
	en.GET("/fonts/:file", func(c *gin.Context) {
		c.File("web/fonts/" + c.Params.ByName("file"))
	})
	en.GET("/js/:file", func(c *gin.Context) {
		c.File("web/js/" + c.Params.ByName("file"))
	})
	en.Any("/favicon.ico", func(c *gin.Context) {
		c.File("web/favicon.ico")
	})
	apiGroup := en.Group("/api/")
	apiGroup.Use(errHandler)
	//redisGroup
	redisApiGroup := apiGroup.Group("redis")
	redisApiGroup.Use(authHander)
	redisApiGroup.Any("/connection/list", connectionController.list)
	redisApiGroup.Any("/connection/info", connectionController.Info)
	redisApiGroup.Any("/connection/test", connectionController.TestConn)
	redisApiGroup.Any("/connection/update", connectionController.Update)
	redisApiGroup.Any("/connection/delete", connectionController.Delete)
	redisApiGroup.Any("/db/list", redisController.Dbs)
	redisApiGroup.Any("/key/list", redisController.KeyList)
	redisApiGroup.Any("/key/update", redisController.UpdateData)
	redisApiGroup.Any("/key/info", redisController.GetKeyInfo)
	redisApiGroup.Any("/key/delete", redisController.Delete)
	redisApiGroup.Any("/info", redisController.Info)

	userApiGroup := apiGroup.Group("user")
	userApiGroup.Use(authHander)
	userApiGroup.Any("/info", userController.Info)
	userApiGroup.Any("/update", userController.Update)

	// 登录登出
	apiGroup.Any("/login", loginController.Login)
	apiGroup.Any("/logout", loginController.Logout)
}
