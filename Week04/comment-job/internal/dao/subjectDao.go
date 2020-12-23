package data

import (
	"fmt"
	"github.com/kaige0207/Go-000/Week04/comment-job/internal/data"
	"github.com/kaige0207/Go-000/Week04/comment-job/internal/pkg/mysqldb"
	"github.com/pkg/errors"
	"log"
	"time"
)

type SubjectDao struct {
	subject data.Subject
}

func (c *SubjectDao) insertSubject(sub *data.Subject) error {
	db, err := mysqldb.NewDB()
	if err != nil {
		log.Println("failed to open database:", err.Error())
		panic(err)
		//return errors.Wrap(err, fmt.Sprintf("dao error: failed to open database: %+v", err.Error()))
	}

	sql := `INSERT INTO comment_subject(	
				obj_id,
				obj_type,
				member_id,
				count,
				root_count,
				all_count,
				state,
				attrs,
				create_time,
				update_time)
            VALUES (?,?,?,?,?,?,?,?,?,?)`
	_, err = db.Exec(
		sql,
		sub.ObjId,
		sub.ObjType,
		sub.MemberId,
		sub.Count,
		sub.RootCount,
		sub.AllCount,
		sub.State,
		sub.Attrs,
		time.Now(),
		time.Now())
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("dao error: insert data failed: %+v", err.Error()))
	}

	return nil
}
