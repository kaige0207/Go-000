package data

import (
	"fmt"
	"github.com/kaige0207/Go-000/Week04/comment-job/internal/data"
	"github.com/kaige0207/Go-000/Week04/comment-job/internal/pkg/mysqldb"
	"github.com/pkg/errors"
	"log"
	"time"
)

type IndexDao struct {
	index data.Index
}

func (c *IndexDao) insertIndex(ind *data.Index) error {
	db, err := mysqldb.NewDB()
	if err != nil {
		log.Println("failed to open database:", err.Error())
		panic(err)
		//return errors.Wrap(err, fmt.Sprintf("dao error: failed to open database: %+v", err.Error()))
	}

	sql := `INSERT INTO comment_index(	
				obj_id,
				obj_type,
				member_id,
				root,
				parent,
				floor,
				count,
				root_count,
				like,
				hate,
				state,
				attrs,
				create_time,
				update_time)
            VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
	_, err = db.Exec(
		sql,
		ind.ObjId,
		ind.ObjType,
		ind.MemberId,
		ind.Root,
		ind.Parent,
		ind.Floor,
		ind.Count,
		ind.RootCount,
		ind.Like,
		ind.Hate,
		ind.State,
		ind.Attrs,
		time.Now(),
		time.Now())
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("dao error: insert data failed: %+v", err.Error()))
	}

	return nil
}
