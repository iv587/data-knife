package api

import "github.com/gin-gonic/gin"

type JsonRes struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type BaseController struct {
}

func successRes(ctx *gin.Context, msg string, data interface{}) {
	res := JsonRes{
		Code: 1,
		Msg:  msg,
		Data: data,
	}
	ctx.JSON(200, res)
}

func errorRes(ctx *gin.Context, msg string, err error) {
	res := JsonRes{
		Code: 0,
		Msg:  msg,
	}
	ctx.JSON(200, res)
}

func notLoginRes(ctx *gin.Context, msg string) {
	res := JsonRes{
		Code: 3,
		Msg:  msg,
	}
	ctx.JSON(200, res)
}
