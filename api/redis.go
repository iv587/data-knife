package api

import (
	"dk/redis"
	"github.com/gin-gonic/gin"
	"strconv"
)

type RedisController struct {
	BaseController
}

var redisController = new(RedisController)

func (r *RedisController) Info(c *gin.Context) {
	var param redis.KeyParam
	if err := c.ShouldBind(&param); err != nil {
		panic(err)
	}
	info, err := redis.Info(param.Id)
	if err != nil {
		panic(err)
	}
	successRes(c, "", info)
}

// Dbs 获取redis数据库列表
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
	successRes(ctx, "", gin.H{
		"list": dbList,
	})
}

// Keys 获取redis的key列表
func (r *RedisController) Keys(ctx *gin.Context) {
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
	successRes(ctx, "", &keyListInfo)
}

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
	successRes(ctx, res, nil)
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
	successRes(c, "", *res)
}
