package dao

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
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

	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Println("connect redis error :", err.Error())
		panic(err)
	}
	defer conn.Close()
	_, err = conn.Do("hset", con.CommentId, "AtMemberIds", con.AtMemberIds, "ip", con.Ip, "Platform", con.Platform, "Device", con.Device, "Message", con.Message, "Meta", con.Meta)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("dao error: insert data failed: %+v", err.Error()))
	}
	return nil
}
