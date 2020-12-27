package dao

import (
	"github.com/kaige0207/Go-000/Week04/comment/internal/data"
	"testing"
	"time"
)

func TestSendComment(t *testing.T) {

	comment := data.Comment{
		Platform:    2,
		CommentId:   0,
		AtMemberIds: "",
		Ip:          0,
		Device:      "",
		Message:     "this is a test massage-1",
		Meta:        "",
		Floor:       0,
		Count:       0,
		RootCount:   0,
		Like:        0,
		Hate:        0,
		CreateTime:  time.Time{},
		UpdateTime:  time.Time{},
	}

	var dao = &CommentDao{comment}
	_ = dao.sendComment(&dao.comment)
}
