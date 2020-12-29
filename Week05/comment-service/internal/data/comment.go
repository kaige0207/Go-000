package data

import "time"

// Comment　用来封装评论在Kafka中传递的消息实体
type Comment struct {
	Platform    int32     `平台:platform`
	CommentId   int64     `主键:comment_id`
	AtMemberIds string    `at列表:at_member_ids`
	Ip          int64     `ip:ip`
	Device      string    `设备:device`
	Message     string    `评论内容:message`
	Meta        string    `评论元数据：背景、字体:meta`
	Floor       int32     `评论楼层:floor`
	Count       int32     `评论总数:count`
	RootCount   int32     `根评论总数:root_count`
	Like        int32     `点赞数:like`
	Hate        int32     `点踩数:hate`
	CreateTime  time.Time `创建时间:create_time`
	UpdateTime  time.Time `修改时间:update_time`
}
