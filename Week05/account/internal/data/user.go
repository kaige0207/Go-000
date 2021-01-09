package data

import "time"

// User 是一个数据库实体对象类型
type User struct {
	Id         uint
	Username   string
	Password   string
	CreateTime time.Time `创建时间:create_time`
	UpdateTime time.Time `修改时间:update_time`
}
