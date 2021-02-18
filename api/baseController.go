package api

import "github.com/gin-gonic/gin"

type jsonRes struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type BaseController struct {
}

func succResJson(ctx *gin.Context, msg string, data interface{}) {
	res := jsonRes{
		Code: 1,
		Msg:  msg,
		Data: data,
	}
	ctx.JSON(200, res)
}

func errResJson(ctx *gin.Context, msg string, err error) {
	res := jsonRes{
		Code: 0,
		Msg:  msg,
	}
	ctx.JSON(200, res)
}

func noLoginJson(ctx *gin.Context) {
	res := jsonRes{
		Code: 3,
		Msg:  "用户登录状态失效",
	}
	ctx.JSON(200, res)
}

func noLoginWithMsg(ctx *gin.Context, msg string) {
	res := jsonRes{
		Code: 3,
		Msg:  msg,
	}
	ctx.JSON(200, res)
}
