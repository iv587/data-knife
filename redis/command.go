package redis

import (
	"dk/connection"
	"dk/ledis"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	StringType = "string"
	ListType   = "list"
	SetType    = "set"
	HashType   = "hash"
	ZsetType   = "zset"
	NoneType   = "none"
)

func Info(id int) ([]*InfoSection, error) {
	rdc, err := GetClient(id, 0)
	defer rdc.Close()
	if err != nil {
		return nil, err
	}
	res, err := rdc.Info()
	if err != nil {
		return nil, err
	}
	lineStrs := strings.Split(res, "\n")
	infoWrapList := make([]*InfoSection, 0, 20)
	var pInfoWrap *InfoSection
	for _, line := range lineStrs {
		if strings.Index(line, "#") == 0 {
			infoWrap := InfoSection{
				Label: line,
				Items: make([]InfoItem, 0, 20),
			}
			pInfoWrap = &infoWrap
			infoWrapList = append(infoWrapList, pInfoWrap)
		} else if line != "" {
			kv := strings.Split(line, ":")
			if len(kv) > 1 {
				pInfoWrap.Items = append(pInfoWrap.Items, InfoItem{
					Name: kv[0], Value: kv[1],
				})
			}
		}
	}
	return infoWrapList, nil
}

// TestConn 测试连接
func TestConn(connection connection.Connection) (bool, error) {
	rdb, err := ledis.Create(connection.Addr, connection.Password, 0)
	defer rdb.Close()
	if err != nil {
		return false, err
	}
	_, err = rdb.Ping()
	if err != nil {
		return false, err
	}
	return true, nil
}

func GetClient(id, dbNo int) (ledis.Cli, error) {
	conn, err := connection.GetById(id)
	if err != nil {
		return nil, errors.New("连接redis失败")
	}
	return ledis.Create(conn.Addr, conn.Password, dbNo)
}

func Dbs(id int) (int, error) {
	rdb, err := GetClient(id, 0)
	defer rdb.Close()

	if err != nil {
		return 0, err
	}
	res, err := rdb.ConfigGet("databases")
	if err != nil {
		return 0, err
	}
	if len(res) > 1 {
		dbNumStr := res[1]
		dbNum, err := strconv.Atoi(dbNumStr)
		if err != nil {
			return 0, nil
		}
		return dbNum, nil
	}
	return 0, errors.New("获取数据库失败")
}

func Keys(id, dbNo int, keyPattern string) (*KeyList, error) {
	rdb, err := GetClient(id, dbNo)
	defer rdb.Close()
	dbsize, err := rdb.DBSize()
	if err != nil {
		return nil, err
	}
	var maxKeys int64 = 1000
	keyList, err := getKeyListByEval(rdb, maxKeys, keyPattern)
	if err != nil {
		if strings.Contains(err.Error(), "command eval not support") {
			keyList, err = getKeyList(rdb, keyPattern, dbsize, maxKeys)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}
	return &KeyList{
		Total: dbsize,
		List:  keyList,
		Max:   maxKeys,
	}, nil
}

func getKeyListByEval(rdb ledis.Cli, maxKeys int64, keyPattern string) ([]KeyInfo, error) {
	var args = make([]interface{}, 0, 20)
	args = append(args, maxKeys)
	if keyPattern != "" {
		args = append(args, keyPattern)
	}
	var keyInfoList interface{}
	keyInfoList, err := rdb.Eval(keysInfoLuaScript, []string{}, args...)
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
		var t = keyInfo[0].([]uint8)
		keyList = append(keyList, KeyInfo{
			Key:  string(t),
			TTL:  keyInfo[1],
			Type: keyInfo[2],
		})
	}
	return keyList, nil
}

func getKeyList(rdb ledis.Cli, keyPattern string, dbSize, maxKeys int64) ([]KeyInfo, error) {
	var err error
	var keys []string
	if dbSize <= maxKeys {
		pattern := "*"
		if keyPattern != "" {
			pattern = keyPattern
		}
		keys, err = rdb.Keys(pattern)
		if err != nil {
			return nil, err
		}
	} else {
		keys = make([]string, 0, maxKeys)
		if keyPattern != "" {
			var cursor int64 = 0
			for {
				var sKeys []string
				sKeys, cursor, err = rdb.Scan(keyPattern, cursor, maxKeys)
				if err != nil {
					return nil, err
				}
				if len(sKeys) > 0 {
					keys = append(keys, sKeys...)
				}
				if cursor == 0 {
					break
				}
			}
		} else {
			for i := 0; i < int(maxKeys); i++ {
				var key string
				key, err = rdb.RandomKey()
				if err != nil {
					return nil, err
				}
				keys = append(keys, key)
			}
		}
	}
	keyInfoList := make([]KeyInfo, 0, len(keys))
	for _, key := range keys {
		var keyType string
		keyType, err = rdb.Type(key)
		if err != nil {
			return nil, err
		}
		ttl, err := rdb.TTL(key)
		if err != nil {
			return nil, err
		}
		keyInfoList = append(keyInfoList, KeyInfo{
			Key:  key,
			Type: keyType,
			TTL:  ttl,
		})
	}
	return keyInfoList, nil
}

// Update 更新 redis 数据
func Update(param KeyParam) (string, error) {
	var (
		res string
		err error
	)
	rdc, err := GetClient(param.Id, param.DbNo)
	defer rdc.Close()

	if err != nil {
		return "", err
	}
	res, err = rdc.Type(param.Key)
	if err != nil {
		return "", err
	}
	if res != NoneType && res != param.Type {
		return "", errors.New("数据类型选择错误")
	}
	if !strings.EqualFold(param.ValuesStr, "") {
		err = json.Unmarshal([]byte(param.ValuesStr), &param.Values)
		if err != nil {
			return "", err
		}
	}
	if !strings.EqualFold(param.DelValuesStr, "") {
		err = json.Unmarshal([]byte(param.DelValuesStr), &param.DelValues)
		if err != nil {
			return "", err
		}
	}
	var ttl int64
	switch param.Type {
	case StringType:
		if param.IsUpdateTtl == 1 {
			ttl = param.Ttl
		} else {
			ttl, err = rdc.TTL(param.Key)
			if err != nil {
				return "", err
			}
			if ttl < 0 {
				ttl = 0
			}
		}
		res, err = updateStrTypeVal(rdc, param.Key, param.Value, ttl)
	case HashType:
		res, err = updateHashTypeVal(rdc, param)
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
func updateStrTypeVal(rdc ledis.Cli, key, val string, ttl int64) (string, error) {
	res, err := rdc.Set(key, val, ttl)
	return res, err
}

// 更新hash类型的数据
func updateHashTypeVal(rdc ledis.Cli, param KeyParam) (string, error) {
	var res string
	var err error
	values, delValues, key := param.Values, param.DelValues, param.Key
	if len(values) > 0 {
		mapVal := make(map[string]string)
		for _, v := range values {
			mapVal[v.Field] = v.Val
		}
		res, err = rdc.HMSet(key, mapVal)
		if err != nil {
			return "", err
		}
	}
	if len(delValues) > 0 {
		res1, err := rdc.HDel(key, delValues)
		if err != nil {
			return "", err
		}
		res = strconv.FormatInt(res1, 10)
	}
	if param.IsUpdateTtl > 0 && param.Ttl > 0 {
		_, err = rdc.Expire(key, param.Ttl)
		if err != nil {
			return "", err
		}
	}
	return res, nil
}

func updateListTypeVal(rdc ledis.Cli, key string, values []FvPairs, ttl int64) (int64, error) {
	strVal := make([]string, 0, len(values))
	for _, v := range values {
		strVal = append(strVal, v.Val)
	}
	rdc.Del(key)
	res, err := rdc.LPush(key, strVal...)
	if err != nil {
		return 0, err
	}
	if ttl > 0 {
		_, err = rdc.Expire(key, ttl)
		if err != nil {
			return 0, err
		}
	}
	return res, nil
}

func updateSetTypeVal(rdc ledis.Cli, key string, values []FvPairs, ttl int64) (int64, error) {
	strVal := make([]string, 0, len(values))
	for _, v := range values {
		strVal = append(strVal, v.Val)
	}
	rdc.Del(key)
	res, err := rdc.SAdd(key, strVal...)
	if err != nil {
		return 0, err
	}
	if ttl > 0 {
		_, err = rdc.Expire(key, ttl)
		if err != nil {
			return 0, err
		}
	}
	return res, nil
}

func updateZetTypeVal(rdc ledis.Cli, key string, values []FvPairs, ttl int64) (int64, error) {
	_, err := rdc.Del(key)
	if err != nil {

	}
	z := make([]ledis.Z, 0, len(values))
	for _, i2 := range values {
		var score float64
		if i2.Score != "" {
			score, err = strconv.ParseFloat(i2.Score, 10)
		}
		z = append(z, ledis.Z{
			Score:  score,
			Member: i2.Val,
		})
	}
	res, err := rdc.ZAdd(key, z...)
	if err != nil {
		return 0, err
	}
	if ttl > 0 {
		_, err = rdc.Expire(key, ttl)
		if err != nil {
			return 0, err
		}
	}
	return res, err
}

// KeyValInfo 获取key value的信息
func KeyValInfo(param KeyParam) (*KeyDetail, error) {
	rdc, err := GetClient(param.Id, param.DbNo)
	defer rdc.Close()
	if err != nil {
		return nil, err
	}
	dataType, err := rdc.Type(param.Key)
	if err != nil {
		return nil, err
	}
	if dataType == NoneType {
		return nil, errors.New("该key不存在或被删除")
	}
	ttlRes, err := rdc.TTL(param.Key)
	if err != nil {
		return nil, err
	}
	var ttl int64
	if ttlRes > 0 {
		ttl = ttlRes
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
		res, err := rdc.Get(param.Key)
		if err != nil {
			return nil, err
		}
		keyDetail.Value = res
	case ListType:
		res, err := rdc.LRange(param.Key, 0, -1)
		if err != nil {
			return nil, err
		}
		keyDetail.Values = make([]FvPairs, 0, len(res))
		for k, v := range res {
			fv := FvPairs{Field: fmt.Sprintf("%d", k), Val: v}
			keyDetail.Values = append(keyDetail.Values, fv)
		}
	case HashType:
		res, err := rdc.HGetAll(param.Key)
		if err != nil {
			return nil, err
		}
		keyDetail.Values = make([]FvPairs, 0, len(res))
		for k, v := range res {
			fv := FvPairs{Field: k, Val: v}
			keyDetail.Values = append(keyDetail.Values, fv)
		}

	case SetType:
		res, err := rdc.SMembers(param.Key)
		if err != nil {
			return nil, err
		}
		keyDetail.Values = make([]FvPairs, 0, len(res))
		for k, v := range res {
			fv := FvPairs{Field: fmt.Sprintf("%d", k), Val: v}
			keyDetail.Values = append(keyDetail.Values, fv)
		}
	case ZsetType:
		z, err := rdc.ZRange(param.Key, 0, -1, true)
		if err != nil {
			return nil, err
		}
		keyDetail.Values = make([]FvPairs, 0, len(z))
		for k, v := range z {
			fv := FvPairs{
				Field: fmt.Sprintf("%d", k),
				Val:   v.Member, Score: fmt.Sprintf("%f", v.Score)}
			keyDetail.Values = append(keyDetail.Values, fv)
		}
	}
	return &keyDetail, err
}

func Delete(param KeyParam) (int64, error) {
	rdc, err := GetClient(param.Id, param.DbNo)
	defer rdc.Close()
	if err != nil {
		return 0, err
	}
	res, err := rdc.Del(param.Key)
	return res, err
}
