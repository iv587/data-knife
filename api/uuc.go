package api

import (
	"dk/auth"
	"dk/user"
	"github.com/gin-gonic/gin"
)

type UucController struct {
	BaseController
}

var uucController = new(UucController)

func (c *UucController) Login(ctx *gin.Context) {
	var param user.User
	if err := ctx.ShouldBind(&param); err != nil {
		panic(err)
	}
	findUser, err := user.GetUserByNameAndPass(param.UserName, param.Password)
	if err != nil {
		panic(err)
	}
	token, err := auth.Token(*findUser)
	if err != nil {
		panic(err)
	}
	successRes(ctx, "", gin.H{
		"token":    token,
		"userName": findUser.UserName,
	})
}

func (c *UucController) appCheck(ctx *gin.Context) {
	count, err := user.CountAll()
	if err != nil {
		panic(err)
	}
	successRes(ctx, "", gin.H{
		"init": count,
	})
}

func (c *UucController) appInit(ctx *gin.Context) {
	var param user.User
	if err := ctx.ShouldBind(&param); err != nil {
		panic(err)
	}
	err := user.InitUser(param)
	if err != nil {
		panic(err)
	}
	successRes(ctx, "初始化成功", gin.H{})
}
