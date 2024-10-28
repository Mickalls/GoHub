package user

import "GoHub/app/models"

type User struct {
	models.BaseModel

	Name     string `json:"name,omitempty"`
	Email    string `json:"-"` // json:"-" 表示令JSON解析器忽略字段,后面接口返回给用户数据的时候,这三个字段会被隐藏
	Phone    string `json:"-"`
	Password string `json:"-"`

	models.CommonTimestampsField
}
