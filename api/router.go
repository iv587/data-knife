package api

import (
	"dk/auth"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/fs"
	"net/http"
	"path"
	"strings"
)

var deviceList = []string{
	"phone", "pad", "pod", "iPhone", "iPod", "ios",
	"iPad", "Android", "Mobile", "BlackBerry", "IEMobile",
	"MQQBrowser", "JUC", "Fennec", "wOSBrowser", "BrowserNG", "WebOS", "Symbian", "Windows Phone"}

func toHTTPError(err error) (msg string, httpStatus int) {
	if errors.Is(err, fs.ErrNotExist) {
		return "404 page not found", http.StatusNotFound
	}
	if errors.Is(err, fs.ErrPermission) {
		return "403 Forbidden", http.StatusForbidden
	}
	// Default:
	return "500 Internal Server Error", http.StatusInternalServerError
}

var Router = func(route *gin.Engine) {
	route.Use(func(context *gin.Context) {
		url := context.Request.URL.Path
		if strings.Index(url, "/api/") != 0 {
			filePath := url
			if strings.EqualFold(url, "/") {
				filePath = "index.html"
				ua := context.Request.UserAgent()
				for _, d := range deviceList {
					if strings.Contains(ua, d) {
						filePath = path.Join("mobile", "index.html")
						break
					}
				}

			} else if strings.EqualFold(url, "/login/") || strings.EqualFold(url, "/login") {
				filePath = path.Join("login", "index.html")
			}
			context.File(path.Join("web", filePath))
			context.Abort()
		} else {
			context.Next()
		}
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
	println("跨域")
	cors.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "all")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "x-requested-with")
		c.Next()
	})
}

var errHandler = func(ctx *gin.Context) {
	println("全局")
	defer func() {
		if err := recover(); err != nil {
			err1 := err.(error)
			fmt.Println(err)
			errorRes(ctx, err1.Error(), err1)
		}
		ctx.Abort()
	}()
	ctx.Next()
}

var authHandler = func(ctx *gin.Context) {
	tokenStr := ctx.PostForm("token")
	ok, user, _ := auth.Verify(tokenStr)
	if !ok {
		notLoginRes(ctx, "")
		ctx.Abort()
	} else {
		ctx.Set("userId", user.Id)
		ctx.Next()
	}
}
