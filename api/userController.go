package api

import (
	"github.com/gin-gonic/gin"
	"github.io/iv587/goredis-admin/db"
)

type UserController struct {
	BaseController
}

var userController = new(UserController)

func (r *UserController) Info(c *gin.Context) {
	var param db.User
	if err := c.ShouldBind(&param); err != nil {
		panic(err)
	}
	u, err := db.GetUserInfo(c.GetInt("userId"))
	if err != nil {
		panic(err)
	}
	u.Id = 0
	succResJson(c, "", u)
}

func (r *UserController) Update(c *gin.Context) {
	var param db.User
	if err := c.ShouldBind(&param); err != nil {
		panic(err)
	}
	param.Id = c.GetInt("userId")
	err := db.UpdateUser(param)
	if err != nil {
		panic(err)
	}
	succResJson(c, "修改成功", gin.H{})
}
