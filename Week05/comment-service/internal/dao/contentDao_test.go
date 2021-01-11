package dao

import (
	"fmt"
	"testing"
)

func TestGetByCommentId(t *testing.T) {
	var con = &ContentDao{}
	commentId := "CommentId"
	cont, err := con.getByCommentId(commentId)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Printf("Platform=%d, Message=%s", cont.Platform, cont.Message)
}
