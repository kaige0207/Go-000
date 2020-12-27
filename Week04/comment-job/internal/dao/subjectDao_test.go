package dao

import (
	"github.com/kaige0207/Go-000/Week04/comment-job/internal/data"
	"testing"
)

func TestInsertSubject(t *testing.T) {
	sub := data.Subject{
		Id:        0,
		ObjId:     0,
		ObjType:   0,
		MemberId:  0,
		Count:     0,
		RootCount: 0,
		AllCount:  0,
		State:     0,
		Attrs:     0,
	}
	var dao = SubjectDao{sub}
	dao.insertSubject(&dao.subject)
}
