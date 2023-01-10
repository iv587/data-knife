package redis

type InfoItem struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type InfoSection struct {
	Label string     `json:"label"`
	Items []InfoItem `json:"items"`
}

type FvPairs struct {
	Field string `json:"field"`
	Val   string `json:"val"`
	Score string `json:"score"`
}

type KeyParam struct {
	Id           int    `form:"id" json:"id"`
	DbNo         int    `form:"dbNo" json:"dbNo"`
	Key          string `form:"key" json:"key"`
	Value        string `form:"value" json:"value"`
	Type         string `form:"type" json:"type"`
	Ttl          int64  `form:"ttl" json:"ttl"`
	IsUpdateTtl  int64  `form:"isUpdateTtl" json:"isUpdateTtl"`
	Values       []FvPairs
	ValuesStr    string `form:"values"`
	DelValues    []string
	DelValuesStr string `form:"delValues"`
}

type KeyInfo struct {
	Key  interface{} `json:"key"`
	Type interface{} `json:"type"`
	TTL  interface{} `json:"ttl"`
}

type KeyList struct {
	List  []KeyInfo `json:"list"`
	Total int64     `json:"total"`
	Max   int64     `json:"max"`
}

type KeyDetail struct {
	KeyInfo
	Value  string    `json:"value"`
	Values []FvPairs `json:"values"`
}
