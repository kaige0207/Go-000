package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"log"
)

/*
	题目：我们在数据库操作的时候，比如dao 层中当遇到一个sql.ErrNoRows的时候，是否应该 Wrap 这个error，抛给上层。为什么，应该怎么做请写出代码？
	回答：应该和其它错误一样 Wrap 向上抛给调用者。这样上层业务可以得到堆栈信息，然后利用 errors.Is()方法判断是否是sql.ErrNoRows错误，进行处理
	助教批改意见：直接将原始的sql错误上报，上下层有依赖哦
*/

type User struct {
	Id   uint
	Name string
}

type ErrUserNotFound struct {
	code uint
	msg string
}

func (err *ErrUserNotFound) Error() string {
	return err.msg
}


//执行数据库查询的函数
func getUserFromDb() (*User, error) {
	//TODO
	//省略数据库逻辑，直接返回sql.ErrNoRows
	return &User{}, sql.ErrNoRows
}

//dao层处理逻辑
func daoFindUserById(uid uint) (*User, error) {
	user, err := getUserFromDb()

	if err != nil {
		if errors.Is(err,sql.ErrNoRows) {
			err = &ErrUserNotFound{404,"this user is not exit"}
		}
		return user, errors.Wrap(err, fmt.Sprintf("dao error: find user by id=%+v", uid))
	}
	return user, nil
}

//biz层处理逻辑
func bizFindUserById(uid uint) (*User, error) {
	return daoFindUserById(uid)
}

func main() {
	user, err := bizFindUserById(1234567)
	if err != nil {
		log.Printf("query user failed: %+v\n", err)
		return
	}

	log.Printf("query user success: user=%+v\n", user)

}
