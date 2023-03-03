package api

import (
	"dk/connection"
	"dk/redis"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ConnectionController struct {
	BaseController
}

var connectionController = new(ConnectionController)

// 获取Redis服务器列表
func (r *ConnectionController) list(ctx *gin.Context) {
	list, err := connection.List()
	fmt.Println(list, err)
	if err != nil {
		println("pa")
		panic(err)
	}
	successRes(ctx, "", gin.H{
		"list": list,
	})
}

func (r *ConnectionController) Info(c *gin.Context) {
	idStr := c.PostForm("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		panic(err)
	}
	data, err := connection.GetById(id)
	if err != nil {
		panic(err)
	}
	successRes(c, "", data)
}

func (r *ConnectionController) Update(c *gin.Context) {
	var conn connection.Connection
	err := c.ShouldBind(&conn)
	if err != nil {
		panic(err)
	}
	upPassStr := c.DefaultPostForm("upPass", "0")
	upPass, err := strconv.Atoi(upPassStr)
	if err != nil {
		panic(err)
	}
	if conn.Id > 0 && upPass == 0 {
		conn.Password = ""
	}
	err = connection.AddOrUpdate(conn)
	if err != nil {
		panic(err)
	}
	successRes(c, "保存成功", "")
}

func (r *ConnectionController) TestConn(c *gin.Context) {
	var con connection.Connection
	err := c.ShouldBind(&con)
	if err != nil {
		panic(err)
	}
	println(con.Addr, con.Password)
	bool, err := redis.TestConn(con)
	if err != nil {
		panic(err)
	}
	if !bool {
		panic(errors.New("连接失败"))
	}
	successRes(c, "测试成功", gin.H{})
}

func (r *ConnectionController) Delete(c *gin.Context) {
	/*var connection db.Connection
	err := c.ShouldBind(&connection)
	if err != nil {
		panic(err)
	}
	err = db.Delete(connection.Id)
	if err != nil {
		panic(err)
	}
	successRes(c, "删除成功", gin.H{})*/
}
