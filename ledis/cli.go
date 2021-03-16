package ledis

type Cli interface {
	Ping() (string, error)

	Info() (string, error)

	Type(key string) (string, error)

	TTL(key string) (int64, error)

	ConfigGet(key string) ([]string, error)

	DBSize() (int64, error)

	Del(key ...string) (int64, error)

	RandomKey() (string, error)

	Eval(lua string, keys []string, args ...interface{}) (interface{}, error)

	Keys(pattern string) ([]string, error)

	Scan(pattern string, cursor, count int64) ([]string, int64, error)

	Expire(key string, sec int64) (int64, error)

	Set(key, val string, tll int64) (string, error)

	Get(key string) (string, error)

	LRange(key string, start, end int) ([]string, error)

	LPush(key string, val ...string) (int64, error)

	HGetAll(key string) (map[string]string, error)

	HMSet(key string, args map[string]string) (string, error)

	SMembers(key string) ([]string, error)

	SAdd(key string, val ...string) (int64, error)

	ZAdd(key string, do ...Z) (int64, error)

	ZRange(key string, start, end int, withScore bool) ([]Z, error)

	Close() error
}

type Z struct {
	Score  float64
	Member string
}
