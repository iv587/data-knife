package ledis

import (
	"github.com/gomodule/redigo/redis"
	"strconv"
)

type RedigoCli struct {
	conn redis.Conn
}

func Create(addr, password string, dbNo int) (*RedigoCli, error) {
	conn, err := redis.Dial("tcp", addr, redis.DialPassword(password), redis.DialDatabase(dbNo))
	if err != nil {
		return nil, err
	}
	rdc := RedigoCli{
		conn: conn,
	}
	return &rdc, nil
}

func (c *RedigoCli) Ping() (string, error) {
	return redis.String(c.conn.Do("PING"))
}

func (c *RedigoCli) Info() (string, error) {
	return redis.String(c.conn.Do("INFO"))
}

func (c *RedigoCli) Type(key string) (string, error) {
	return redis.String(c.conn.Do("TYPE", key))
}

func (c *RedigoCli) TTL(key string) (int64, error) {
	return redis.Int64(c.conn.Do("TTL", key))
}

func (c *RedigoCli) ConfigGet(key string) ([]string, error) {
	return redis.Strings(c.conn.Do("CONFIG", "GET", key))
}

func (c *RedigoCli) DBSize() (int64, error) {
	return redis.Int64(c.conn.Do("DBSIZE"))
}

func (c *RedigoCli) Del(key ...string) (int64, error) {
	if keylen := len(key); keylen > 0 {
		args := make([]interface{}, 0, keylen)
		for _, i2 := range key {
			args = append(args, i2)
		}
		return redis.Int64(c.conn.Do("DEL", args...))
	}
	return 0, nil
}

func (c *RedigoCli) RandomKey() (string, error) {
	return redis.String(c.conn.Do("RANDOMKEY"))
}

func (c *RedigoCli) Eval(lua string, keys []string, args ...interface{}) (interface{}, error) {
	cmdArgs := make([]interface{}, 0, len(keys)+len(args)+2)
	cmdArgs = append(cmdArgs, lua)
	cmdArgs = append(cmdArgs, len(keys))
	for _, i2 := range keys {
		cmdArgs = append(cmdArgs, i2)
	}
	for _, i2 := range args {
		cmdArgs = append(cmdArgs, i2)
	}
	return c.conn.Do("EVAL", cmdArgs...)
}

func (c *RedigoCli) Keys(pattern string) ([]string, error) {
	return redis.Strings(c.conn.Do("KEYS", pattern))
}

func (c *RedigoCli) Scan(pattern string, cursor, count int64) ([]string, int64, error) {
	args := make([]interface{}, 0, 5)
	args = append(args, cursor)
	if pattern != "" {
		args = append(args, "MATCH", pattern)
	}
	if count > 0 {
		args = append(args, "COUNT", count)
	}
	//res, err := c.conn.Do("SCAN", args...)
	res, err := c.conn.Do("SCAN", args...)
	if err != nil {
		return nil, 0, err
	}
	ress := res.([]interface{})
	nextCorStr := string(ress[0].([]uint8))
	nextCor, err := strconv.ParseInt(nextCorStr, 10, 64)
	if err != nil {
		return nil, 0, err
	}
	valS := ress[1].([]interface{})
	var values []string
	if len(valS) > 0 {
		values = make([]string, 0, len(valS))
		for _, i2 := range valS {
			values = append(values, string(i2.([]uint8)))
		}
	}
	return values, nextCor, nil
}

func (c *RedigoCli) Expire(key string, sec int64) (int64, error) {
	return redis.Int64(c.conn.Do("EXPIRE", key, sec))
}

func (c *RedigoCli) Set(key, val string, tll int64) (string, error) {
	re, err := redis.String(c.conn.Do("SET", key, val))
	if err != nil {
		return "", err
	}
	if tll > 0 {
		_, err = c.Expire(key, tll)
		if err != nil {
			return "", err
		}
	}
	return re, nil
}

func (c *RedigoCli) Get(key string) (string, error) {
	return redis.String(c.conn.Do("GET", key))
}

func (c *RedigoCli) LRange(key string, start, end int) ([]string, error) {
	return redis.Strings(c.conn.Do("LRANGE", key, start, end))
}

func (c *RedigoCli) LPush(key string, val ...string) (int64, error) {
	if argLen := len(val); argLen > 0 {
		args := make([]interface{}, 0, argLen+1)
		args = append(args, key)
		for _, i2 := range val {
			args = append(args, i2)
		}
		return redis.Int64(c.conn.Do("LPUSH", args...))
	}
	return 0, nil
}

func (c *RedigoCli) HGetAll(key string) (map[string]string, error) {
	res, err := redis.Strings(c.conn.Do("HGETALL", key))
	if err != nil {
		return nil, err
	}
	var resMap map[string]string
	if len(res) > 0 {
		resMap = make(map[string]string)
		for i := 0; i < len(res); i = i + 2 {
			resMap[res[i]] = res[i+1]
		}
	}
	return resMap, nil
}

func (c *RedigoCli) HMSet(key string, args map[string]string) (string, error) {
	if argLen := len(args); argLen > 0 {
		argArray := make([]interface{}, 0, argLen*2+1)
		argArray = append(argArray, key)
		for i, i2 := range args {
			argArray = append(argArray, i, i2)
		}
		return redis.String(c.conn.Do("HMSET", argArray...))
	}
	return "", nil
}

func (c *RedigoCli) HDel(key string, args []string) (int64, error) {
	if argLen := len(args); argLen > 0 {
		argArray := make([]interface{}, 0, argLen*2+1)
		argArray = append(argArray, key)
		for i, i2 := range args {
			argArray = append(argArray, i, i2)
		}
		return redis.Int64(c.conn.Do("HDel", argArray...))
	}
	return 0, nil
}

func (c *RedigoCli) SMembers(key string) ([]string, error) {
	return redis.Strings(c.conn.Do("SMEMBERS", key))
}

func (c *RedigoCli) SAdd(key string, val ...string) (int64, error) {
	if argLen := len(val); argLen > 0 {
		args := make([]interface{}, 0, argLen+1)
		args = append(args, key)
		for _, i2 := range val {
			args = append(args, i2)
		}
		return redis.Int64(c.conn.Do("SADD", args...))
	}
	return 0, nil
}

func (c *RedigoCli) ZAdd(key string, do ...Z) (int64, error) {
	if zLen := len(do); zLen > 0 {
		args := make([]interface{}, 0, zLen*2+1)
		args = append(args, key)
		for _, i2 := range do {
			args = append(args, i2.Score, i2.Member)
		}
		return redis.Int64(c.conn.Do("ZADD", args...))
	}
	return 0, nil
}

func (c *RedigoCli) ZRange(key string, start, end int, withScore bool) ([]Z, error) {
	args := make([]interface{}, 0, 4)
	args = append(args, key, start, end)
	if withScore {
		args = append(args, "WITHSCORES")
	}
	res, err := c.conn.Do("ZRANGE", args...)
	resInt := res.([]interface{})
	values := make([]Z, 0, len(resInt))
	for i := 0; i < len(resInt); i = i + 2 {
		score, err := strconv.ParseFloat(string(resInt[i+1].([]uint8)), 64)
		if err != nil {
			return nil, err
		}
		values = append(values, Z{
			Member: string(resInt[i].([]uint8)),
			Score:  score,
		})

	}
	if err != nil {
		return nil, err
	}
	return values, nil
}

func (c *RedigoCli) Close() error {
	return c.conn.Close()
}
