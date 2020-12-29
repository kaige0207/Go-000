package data

import "time"

// Subject 类型对应 MySQL数据库中的 comment_subject 数据表
type Subject struct {
	Id         int64     `主键:id`
	ObjId      int64     `对象id:obj_id`
	ObjType    int8      `对象类型:obj_type`
	MemberId   int64     `发表者用户id:member_id`
	Count      int32     `评论总数:count`
	RootCount  int32     `根评论总数:root_count`
	AllCount   int32     `评论+回复总数:all_count`
	State      int8      `状态(0-正常、1-隐藏):state`
	Attrs      int32     `属性(0-运营置、1-up置顶、2-大数据过滤):attrs`
	CreateTime time.Time `创建时间:create_time`
	UpdateTime time.Time `修改时间:update_time`
}
