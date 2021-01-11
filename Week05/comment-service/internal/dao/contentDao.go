package dao

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/kaige0207/Go-000/Week05/comment/internal/data"
	"log"
	"strconv"
)

type ContentDao struct {
	content data.Content
}

func (c *ContentDao) getByCommentId(commentId string) (content data.Content, err error) {
	arr := [10]string{}
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Println("connect redis error :", err.Error())
		panic(err)
	}
	defer conn.Close()
	res, err := redis.Values(conn.Do("HMGET", commentId, "Platform", "Message"))
	if err != nil {
		log.Println("redis HGET error:", err)
	} else {
		for i, v := range res {
			fmt.Printf("%s\n", v.([]byte))
			arr[i] = string(v.([]byte))
		}
	}
	platform, _ := strconv.Atoi(arr[0])
	content.Platform = int8(platform)
	content.Message = arr[1]
	return content, err
}
