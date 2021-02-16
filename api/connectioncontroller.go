package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.io/iv587/goredis-admin/db"
	"github.io/iv587/goredis-admin/redis"
	"strconv"
)

type ConnectionController struct {
	BaseController
}

var connectionController = new(ConnectionController)

//获取Redis服务器列表
func (r *ConnectionController) list(ctx *gin.Context) {
	list, err := db.ConnectionList()
	if err != nil {
		panic(err)
	}
	succResJson(ctx, "", gin.H{
		"list": list,
	})
}

func (r *ConnectionController) Info(c *gin.Context) {
	idStr := c.PostForm("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		panic(err)
	}
	data, err := db.GetConnectionById(id)
	if err != nil {
		panic(err)
	}
	succResJson(c, "", data)
}

func (r *ConnectionController) Update(c *gin.Context) {
	var connection db.Connection
	err := c.ShouldBind(&connection)
	if err != nil {
		panic(err)
	}
	upPassStr := c.DefaultPostForm("upPass", "0")
	upPass, err := strconv.Atoi(upPassStr)
	if err != nil {
		panic(err)
	}
	err = redis.AddOrUpdateConn(connection, upPass)
	if err != nil {
		panic(err)
	}
	succResJson(c, "保存成功", "")
}

func (r *ConnectionController) TestConn(c *gin.Context) {
	var connection db.Connection
	err := c.ShouldBind(&connection)
	if err != nil {
		panic(err)
	}
	bool, err := redis.TestConn(connection.Addr, connection.Password)
	if err != nil {
		panic(err)
	}
	if !bool {
		panic(errors.New("连接失败"))
	}
	succResJson(c, "测试成功", gin.H{})
}

func (r *ConnectionController) Delete(c *gin.Context) {
	var connection db.Connection
	err := c.ShouldBind(&connection)
	if err != nil {
		panic(err)
	}
	err = db.Delete(connection.Id)
	if err != nil {
		panic(err)
	}
	succResJson(c, "删除成功", gin.H{})
}
