package data

import "time"

// Index 类型对应 MySQL数据库中的 comment_index 数据表
type Index struct {
	Id         int64     `主键:id`
	ObjId      int64     `对象id:obj_id`
	ObjType    int8      `对象类型:obj_type`
	MemberId   int64     `发表者用户id:member_id`
	Root       int64     `根评论id，不为0是回复评论:root`
	Parent     int64     `回复评论id，为0是root评论:parent`
	Floor      int32     `评论楼层:floor`
	Count      int32     `评论总数:count`
	RootCount  int32     `根评论总数:root_count`
	Like       int32     `点赞数:like`
	Hate       int32     `点踩数:hate`
	State      int8      `状态(0-正常、1-隐藏):state`
	Attrs      int32     `属性(0-运营置、1-up置顶、2-大数据过滤):attrs`
	CreateTime time.Time `创建时间:create_time`
	UpdateTime time.Time `修改时间:update_time`
}
