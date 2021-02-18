package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.io/iv587/goredis-admin/db"
	"strconv"
	"strings"
	"time"
)

const (
	StringType = "string"
	ListType   = "list"
	SetType    = "set"
	HashType   = "hash"
	ZsetType   = "zset"
	NoneType   = "none"
)

// 添加连接
func AddOrUpdateConn(conn db.Connection, upPass int) error {
	if conn.Id <= 0 || upPass == 1 {
		isAvaliable, err := TestConn(conn.Addr, conn.Password)
		if !isAvaliable {
			return err
		}
	}
	//添加文件
	err := db.UpdateConnection(conn, upPass)
	return err
}

// 测试连接
func TestConn(addr, password string) (bool, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:        addr,
		Password:    password,
		IdleTimeout: time.Minute * 5,
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return false, err
	}
	return true, nil
}

func Dbs(id int) (int, error) {
	rdb, err := db.GetClient(id, 0)
	if err != nil {
		return 0, err
	}
	res, err := rdb.ConfigGet(context.Background(), "databases").Result()
	if err != nil {
		return 0, err
	}
	if len(res) > 1 {
		dbNumStr, _ := res[1].(string)
		dbNum, err := strconv.Atoi(dbNumStr)
		if err != nil {
			return 0, nil
		}
		return dbNum, nil
	}
	return 0, errors.New("获取数据库失败")
}

func Keys(id, dbNo int, keyPattern string) (*KeyList, error) {
	rdb, err := db.GetClient(id, dbNo)
	if err != nil {
		return nil, err
	}
	var args = make([]interface{}, 0, 20)
	args = append(args, 100)
	if keyPattern != "" {
		args = append(args, keyPattern)
	}
	var keyInfoList interface{}
	keyInfoList, err = rdb.Eval(context.Background(), keysInfoLuaScript, []string{}, args).Result()
	if err != nil {
		return nil, err
	}
	keyInfoList1, ok := keyInfoList.([]interface{})
	if !ok {
		return nil, errors.New("数据转换失败")
	}
	keyList := make([]KeyInfo, 0, len(keyInfoList1))
	for _, v := range keyInfoList1 {
		keyInfo := v.([]interface{})
		keyList = append(keyList, KeyInfo{
			Key:  keyInfo[0],
			TTL:  keyInfo[1],
			Type: keyInfo[2],
		})
	}
	return &KeyList{
		Total: rdb.DBSize(context.Background()).Val(),
		List:  keyList,
	}, nil
}

// 更新 redis 数据
func Update(param KeyParam) (string, error) {
	var (
		res string
		err error
	)
	rdc, err := db.GetClient(param.Id, param.DbNo)
	if err != nil {
		return "", err
	}
	res, err = rdc.Type(context.Background(), param.Key).Result()
	if err != nil {
		return "", err
	}
	if res != NoneType && res != param.Type {
		return "", errors.New("已存在其他类型的数据")
	}
	var ttl time.Duration
	if param.IsUpdateTtl == 1 {
		ttl = time.Second * time.Duration(param.Ttl)
	} else {
		ttl, err = rdc.TTL(context.Background(), param.Key).Result()
		if err != nil {
			return "", err
		}
		if ttl < 0 {
			ttl = 0 * time.Second
		}
	}
	err = json.Unmarshal([]byte(param.ValuesStr), &param.Values)
	if err != nil {
		return "", err
	}
	switch param.Type {
	case StringType:
		res, err = updateStrTypeVal(rdc, param.Key, param.Value, ttl)
	case HashType:
		res, err = updateHashTypeVal(rdc, param.Key, param.Values, ttl)
	case ListType:
		var res1 int64
		res1, err = updateListTypeVal(rdc, param.Key, param.Values, ttl)
		res = strconv.FormatInt(res1, 10)
	case SetType:
		var res1 int64
		res1, err = updateSetTypeVal(rdc, param.Key, param.Values, ttl)
		res = strconv.FormatInt(res1, 10)
	case ZsetType:
		var res1 int64
		res1, err = updateZetTypeVal(rdc, param.Key, param.Values, ttl)
		res = strconv.FormatInt(res1, 10)
	}
	return res, err
}

// 更新字符串类型的redis值
func updateStrTypeVal(rdc *redis.Client, key, val string, ttl time.Duration) (string, error) {
	res, err := rdc.Set(context.Background(), key, val, ttl).Result()
	return res, err
}

func updateHashTypeVal(rdc *redis.Client, key string, values []FvPairs, ttl time.Duration) (string, error) {
	rdc.Del(context.Background(), key)
	mapVal := make([]string, 0, len(values)*2)
	for _, v := range values {
		mapVal = append(mapVal, v.Field, v.Val)
	}
	res, err := rdc.HMSet(context.Background(), key, mapVal).Result()
	if err != nil {

		return "", err
	}
	if ttl.Microseconds() > 0 {
		_, err = rdc.Expire(context.Background(), key, ttl).Result()
		if err != nil {
			return "", err
		}
	}
	return strconv.FormatBool(res), nil
}

func updateListTypeVal(rdc *redis.Client, key string, values []FvPairs, ttl time.Duration) (int64, error) {
	strVal := make([]string, 0, len(values))
	for _, v := range values {
		strVal = append(strVal, v.Val)
	}
	rdc.Del(context.Background(), key)
	res, err := rdc.LPush(context.Background(), key, strVal).Result()
	if err != nil {
		return 0, err
	}
	if ttl.Microseconds() > 0 {
		_, err = rdc.Expire(context.Background(), key, ttl).Result()
		if err != nil {
			return 0, err
		}
	}
	return res, nil
}

func updateSetTypeVal(rdc *redis.Client, key string, values []FvPairs, ttl time.Duration) (int64, error) {
	strVal := make([]string, 0, len(values))
	for _, v := range values {
		strVal = append(strVal, v.Val)
	}
	rdc.Del(context.Background(), key)
	res, err := rdc.SAdd(context.Background(), key, strVal).Result()
	if err != nil {
		return 0, err
	}
	if ttl.Microseconds() > 0 {
		_, err = rdc.Expire(context.Background(), key, ttl).Result()
		if err != nil {
			return 0, err
		}
	}
	return res, nil
}

func updateZetTypeVal(rdc *redis.Client, key string, values []FvPairs, ttl time.Duration) (int64, error) {
	_, err := rdc.Del(context.Background(), key).Result()
	if err != nil {

	}
	z := make([]*redis.Z, 0, len(values))
	for _, i2 := range values {
		var score float64
		if i2.Score != "" {
			score, err = strconv.ParseFloat(i2.Score, 10)
		}
		z = append(z, &redis.Z{
			Score:  score,
			Member: i2.Val,
		})
	}
	res, err := rdc.ZAdd(context.Background(), key, z...).Result()
	if err != nil {
		return 0, err
	}
	if ttl.Microseconds() > 0 {
		_, err = rdc.Expire(context.Background(), key, ttl).Result()
		if err != nil {
			return 0, err
		}
	}
	return res, err
}

// 获取key value的信息
func KeyValInfo(param KeyParam) (*KeyDetail, error) {
	rdc, err := db.GetClient(param.Id, param.DbNo)
	if err != nil {
		return nil, err
	}
	dataType, err := rdc.Type(context.Background(), param.Key).Result()
	if err != nil {
		return nil, err
	}
	if dataType == NoneType {
		return nil, errors.New("该key不存在或被删除")
	}
	ttlRes, err := rdc.TTL(context.Background(), param.Key).Result()
	if err != nil {
		return nil, err
	}
	ttl := 0
	if ttlRes > 0 {
		ttl = int(ttlRes.Seconds())
	}
	var keyDetail = KeyDetail{
		KeyInfo: KeyInfo{
			Key:  param.Key,
			Type: dataType,
			TTL:  ttl,
		},
	}
	switch dataType {
	case StringType:
		res, err := rdc.Get(context.Background(), param.Key).Result()
		if err != nil {
			return nil, err
		}
		keyDetail.Value = res
	case ListType:
		res, err := rdc.LRange(context.Background(), param.Key, 0, -1).Result()
		if err != nil {
			return nil, err
		}
		keyDetail.Values = make([]FvPairs, 0, len(res))
		for k, v := range res {
			fv := FvPairs{Field: fmt.Sprintf("%d", k), Val: v}
			keyDetail.Values = append(keyDetail.Values, fv)
		}
	case HashType:
		res, err := rdc.HGetAll(context.Background(), param.Key).Result()
		if err != nil {
			return nil, err
		}
		keyDetail.Values = make([]FvPairs, 0, len(res))
		for k, v := range res {
			fv := FvPairs{Field: k, Val: v}
			keyDetail.Values = append(keyDetail.Values, fv)
		}

	case SetType:
		res, err := rdc.SMembers(context.Background(), param.Key).Result()
		if err != nil {
			return nil, err
		}
		keyDetail.Values = make([]FvPairs, 0, len(res))
		for k, v := range res {
			fv := FvPairs{Field: fmt.Sprintf("%d", k), Val: v}
			keyDetail.Values = append(keyDetail.Values, fv)
		}
	case ZsetType:
		z, err := rdc.ZRangeWithScores(context.Background(), param.Key, 0, -1).Result()
		if err != nil {
			return nil, err
		}
		keyDetail.Values = make([]FvPairs, 0, len(z))
		for k, v := range z {
			fv := FvPairs{Field: fmt.Sprintf("%d", k), Val: v.Member.(string), Score: fmt.Sprintf("%f", v.Score)}
			keyDetail.Values = append(keyDetail.Values, fv)
		}
	}
	return &keyDetail, err
}

func Info(id, dbNo int) ([]*InfoWrap, error) {
	rdc, err := db.GetClient(id, dbNo)
	if err != nil {
		return nil, err
	}
	res, err := rdc.Info(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	lineStrs := strings.Split(res, "\n")
	infoWrapList := make([]*InfoWrap, 0, 20)
	var pInfoWrap *InfoWrap
	for _, line := range lineStrs {
		if strings.Index(line, "#") == 0 {
			infoWrap := InfoWrap{
				Section: line,
				Values:  make([]FvPairs, 0, 20),
			}
			pInfoWrap = &infoWrap
			infoWrapList = append(infoWrapList, pInfoWrap)
		} else if line != "" {
			kv := strings.Split(line, ":")
			if len(kv) > 1 {
				pInfoWrap.Values = append(pInfoWrap.Values, FvPairs{
					Field: kv[0], Val: kv[1],
				})
			}
		}
	}
	return infoWrapList, nil
}

// 删除key
func Delete(param KeyParam) (int64, error) {
	rdc, err := db.GetClient(param.Id, param.DbNo)
	if err != nil {
		return 0, err
	}
	res, err := rdc.Del(context.Background(), param.Key).Result()
	return res, err
}
