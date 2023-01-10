package api

import (
	"dk/auth"
	"github.com/gin-gonic/gin"
)

var Router = func(route *gin.Engine) {
	route.GET("/", func(c *gin.Context) {
		c.File("web/index.html")
	})
	route.GET("/assets/:file", func(c *gin.Context) {
		c.File("web/assets/" + c.Params.ByName("file"))
	})
	route.GET("/vite.svg", func(c *gin.Context) {
		c.File("web/" + c.Params.ByName("file"))
	})
	route.NoRoute(func(context *gin.Context) {
		context.File("web/index.html")
	})

	apiGroup := route.Group("/api")
	apiGroup.Use(errHandler)
	apiGroup.Any("/login", uucController.Login)
	apiGroup.Any("/app/check", uucController.appCheck)
	apiGroup.Any("/app/init", uucController.appInit)
	connectionApiGroup := apiGroup.Group("/connection")
	connectionApiGroup.Use(authHandler)
	connectionApiGroup.Any("/list", connectionController.list)
	connectionApiGroup.Any("/test", connectionController.TestConn)
	connectionApiGroup.Any("/update", connectionController.Update)
	connectionApiGroup.Any("/info", connectionController.Info)

	redisApiGroup := apiGroup.Group("/redis")
	redisApiGroup.Use(authHandler)
	redisApiGroup.POST("/info", redisController.Info)
	redisApiGroup.POST("/dbs", redisController.Dbs)
	redisApiGroup.POST("/keys", redisController.Keys)
	redisApiGroup.POST("/key/info", redisController.GetKeyInfo)
	redisApiGroup.POST("/key/update", redisController.UpdateData)

	userApiGroup := apiGroup.Group("/user")
	userApiGroup.Use(authHandler)
	userApiGroup.POST("/update/password", userController.UpdatePassword)

}

var Cors = func(cors *gin.Engine) {
	cors.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "all")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "x-requested-with")
		c.Next()
	})
}

var errHandler = func(ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			err1 := err.(error)
			errorRes(ctx, err1.Error(), err1)
		}
	}()
	ctx.Next()
}

var authHandler = func(ctx *gin.Context) {
	tokenStr := ctx.PostForm("token")
	ok, user, err := auth.Verify(tokenStr)
	if err != nil {
		panic(err)
	}
	if !ok {
		notLoginRes(ctx, "")
		ctx.Abort()
	} else {
		ctx.Set("userId", user.Id)
		ctx.Next()
	}
}
