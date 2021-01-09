package dao

import (
	"github.com/kaige0207/Go-000/Week05/comment-job/internal/data"
	"testing"
)

func TestInsertContent(t *testing.T) {
	con := data.Content{
		CommentId:   0,
		AtMemberIds: "",
		Ip:          0,
		Platform:    0,
		Device:      "",
		Message:     "",
		Meta:        "",
	}
	var dao = ContentDao{con}
	_ = dao.insertContent(&dao.content)
}
