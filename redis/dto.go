package redis

type ConnectionInfo struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Addr     string `json:"addr"`
	Password string `json:"password"`
}

type KeyInfo struct {
	Key  interface{} `json:"key"`
	Type interface{} `json:"type"`
	TTL  interface{} `json:"ttl"`
}

type KeyList struct {
	List  []KeyInfo `json:"list"`
	Total int64     `json:"total"`
}

type KeyParam struct {
	Id          int    `form:"id" json:"id"`
	DbNo        int    `form:"dbNo" json:"dbNo"`
	Key         string `form:"key" json:"key"`
	Value       string `form:"value" json:"value"`
	Type        string `form:"type" json:"type"`
	Ttl         int64  `form:"ttl" json:"ttl"`
	IsUpdateTtl int64  `form:"isUpdateTtl" json:"ttl"`
	Values      []FvPairs
	ValuesStr   string `form:"values"`
}

type KeyDetail struct {
	KeyInfo
	Value  string    `json:"value"`
	Values []FvPairs `json:"values"`
}

type FvPairs struct {
	Field string `json:"field"`
	Val   string `json:"val"`
	Score string `json:"score"`
}

type InfoWrap struct {
	Section string    `json:"section"`
	Values  []FvPairs `json:"values"`
}
