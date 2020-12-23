package data

import "time"

// Content 类型对应MySQL表中的 comment_content 数据表
type Content struct {
	CommentId   int64     `主键:comment_id`
	AtMemberIds string    `at列表:at_member_ids`
	Ip          int64     `ip:ip`
	Platform    int8      `平台:platform`
	Device      string    `设备:device`
	Message     string    `评论内容:message`
	Meta        string    `评论元数据：背景、字体:meta`
	CreateTime  time.Time `创建时间:create_time`
	UpdateTime  time.Time `修改时间:update_time`
}
