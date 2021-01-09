package dao

import (
	"fmt"
	"github.com/kaige0207/Go-000/Week05/comment-job/internal/data"
	"github.com/kaige0207/Go-000/Week05/comment-job/internal/pkg/mysqldb"
	"github.com/pkg/errors"
	"log"
	"time"
)

type ContentDao struct {
	content data.Content
}

func (c *ContentDao) insertContent(con *data.Content) error {
	db, err := mysqldb.NewDB()
	if err != nil {
		log.Println("failed to open database:", err.Error())
		panic(err)
		//return errors.Wrap(err, fmt.Sprintf("dao error: failed to open database: %+v", err.Error()))
	}

	sql := `INSERT INTO comment_content(
				comment_id,
				at_member_ids,
				ip,
				platform,
				device,
				message,
				meta,
				create_time,
				update_time)
            VALUES (?,?,?,?,?,?,?,?,?)`
	_, err = db.Exec(
		sql,
		con.CommentId,
		con.AtMemberIds,
		con.Ip,
		con.Platform,
		con.Device,
		con.Message,
		con.Meta,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("dao error: insert data failed: %+v", err.Error()))
	}

	return nil
}
