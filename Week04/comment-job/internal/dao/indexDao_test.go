package dao

import (
	"github.com/kaige0207/Go-000/Week04/comment-job/internal/data"
	"testing"
)

func TestInsertIndex(t *testing.T) {
	ind := data.Index{
		Id:        0,
		ObjId:     0,
		ObjType:   0,
		MemberId:  0,
		Root:      0,
		Parent:    0,
		Floor:     0,
		Count:     0,
		RootCount: 0,
		Like:      0,
		Hate:      0,
		State:     0,
		Attrs:     0,
	}
	var dao = IndexDao{ind}
	dao.insertIndex(&dao.index)
}
