package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.io/iv587/goredis-admin/redis"
	"strconv"
)

var redisController = new(RedisController)

type RedisController struct {
	BaseController
}

//获取redis数据库列表
func (r *RedisController) Dbs(ctx *gin.Context) {
	idStr := ctx.PostForm("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		panic(err)
	}
	dbNo, err := redis.Dbs(id)
	if err != nil {
		panic(err)
	}
	dbList := make([]int, 0, dbNo)
	for i := 0; i < dbNo; i++ {
		dbList = append(dbList, i)
	}
	succResJson(ctx, "", gin.H{
		"list": dbList,
	})
}

//获取redis的key列表
func (r *RedisController) KeyList(ctx *gin.Context) {
	idStr := ctx.PostForm("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		panic(err)
	}
	dbNoStr := ctx.PostForm("dbNo")
	dbNo, err := strconv.Atoi(dbNoStr)
	if err != nil {
		panic(err)
	}
	keyPattern := ctx.DefaultPostForm("keyPattern", "")
	keyListInfo, err := redis.Keys(id, dbNo, keyPattern)
	if err != nil {
		panic(err)
	}
	succResJson(ctx, "", &keyListInfo)
}

//更新redis数据
func (r *RedisController) UpdateData(ctx *gin.Context) {
	var param redis.KeyParam
	var (
		res string
		err error
	)
	if err = ctx.ShouldBind(&param); err != nil {
		panic(err)
	}
	res, err = redis.Update(param)
	if err != nil {
		panic(err)
	}
	succResJson(ctx, res, nil)
}

func (r *RedisController) GetKeyInfo(c *gin.Context) {
	var param redis.KeyParam
	if err := c.ShouldBind(&param); err != nil {
		panic(err)
	}
	res, err := redis.KeyValInfo(param)
	if err != nil {
		panic(err)
	}
	succResJson(c, "", *res)
}

func (r *RedisController) Info(c *gin.Context) {
	var param redis.KeyParam
	if err := c.ShouldBind(&param); err != nil {
		panic(err)
	}
	info, err := redis.Info(param.Id, param.DbNo)
	if err != nil {
		panic(err)
	}
	succResJson(c, "", info)
}

func (r *RedisController) Delete(c *gin.Context) {
	var param redis.KeyParam
	var (
		res int64
		err error
	)
	if err = c.ShouldBind(&param); err != nil {
		panic(err)
	}
	res, err = redis.Delete(param)
	if err != nil {
		panic(err)
	}
	succResJson(c, fmt.Sprintf("结果为%d", res), gin.H{})
}
