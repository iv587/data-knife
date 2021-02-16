package api

import (
	"github.com/gin-gonic/gin"
	"github.io/iv587/goredis-admin/auth"
	"github.io/iv587/goredis-admin/db"
	"sync"
	"time"
)

type LoginController struct {
	BaseController
}

var loginController = &LoginController{}

type token struct {
	id         int
	createTime time.Duration
}

var tokenMap = struct {
	sync.RWMutex
	token map[string]int
}{
	token: make(map[string]int),
}

func (c *LoginController) Login(ctx *gin.Context) {
	var param db.User
	if err := ctx.ShouldBind(&param); err != nil {
		panic(err)
	}
	user, err := db.GetUserByNameAndPass(param.Name, param.Password)
	if err != nil {
		panic(err)
	}
	token, err := auth.Gen(user.Id)
	if err != nil {
		panic(err)
	}
	succResJson(ctx, "", gin.H{
		"token": token,
		"name":  user.Name,
	})
}

func (c *LoginController) Logout(ctx *gin.Context) {

}
