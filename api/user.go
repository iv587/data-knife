package api

import (
	"dk/user"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	BaseController
}

var userController = new(UserController)

func (r *UserController) UpdatePassword(c *gin.Context) {
	userId := c.GetInt("userId")
	oldPassword := c.PostForm("oldPassword")
	newPassword := c.PostForm("newPassword")
	user.UpdatePasswd(oldPassword, newPassword, userId)
	successRes(c, "修改成功", gin.H{})
}
