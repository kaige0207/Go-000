package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)
/*
	题目：我们在数据库操作的时候，比如dao 层中当遇到一个sql.ErrNoRows的时候，是否应该 Wrap 这个error，抛给上层。为什么，应该怎么做请写出代码？
	回答：应该和其它错误一样 Wrap 向上抛给调用者。这样上层业务可以得到堆栈信息，然后利用 errors.Is()方法判断是否是sql.ErrNoRows错误，进行处理
 */


type User struct {
	Id 	 uint
	Name string
}

//执行数据库查询的函数
func DaogetUserFromDb() (*User,error) {
	//TODO
	//省略数据库逻辑，直接返回sql.ErrNoRows
	return &User{}, sql.ErrNoRows
}

//dao层处理逻辑
func DaoFindUserById(uid uint) (*User, error) {
	user,err := DaogetUserFromDb()
	if err != nil{
		return user,errors.Wrap(err, fmt.Sprintf("dao error : find user by id=%+v", uid))
	}
	return user, nil
}

//biz层处理逻辑
func BizFindUserById(uid uint) (*User,error) {
	return DaoFindUserById(uid)
}

func main() {
	user, err := BizFindUserById(1234567)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Printf("user not exists %+v\n", err)
			return
		}

		fmt.Printf("query user failed: %+v\n", err)
		return
	}

	fmt.Printf("query user success : user=%+v\n", user)

}
