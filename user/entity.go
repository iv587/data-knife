package user

type User struct {
	Id       int    `json:"id" gorm:"primaryKey" form:"id"`
	UserName string `json:"userName" form:"userName"`
	Password string `json:"password" form:"password"`
}

func (User) TableName() string {
	return "user"
}
