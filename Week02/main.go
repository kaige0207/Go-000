package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"log"
)

type User struct {
	Id 	 int
	Name string
}

func (u *User) DbNoDataError() (*User,error) {

	return u,sql.ErrNoRows
}

//dao层处理逻辑
func (u *User) DaoFindUserById(uid int) (*User,error) {
	if _,err := u.DbNoDataError(); err != nil{
		fmt.Errorf("accessing DB: %w", err)
		return nil,err
	}
	return u, nil
}

//biz层处理逻辑
func BizFindUserById(uid int) (*User,error) {
	user := &User{Id:uid}
	user, err := user.DaoFindUserById(uid)
	if err != nil {
		return nil, errors.WithMessagef(err, "biz query user: %d detail", uid)
	}
	return user.DaoFindUserById(uid)
}


func main() {
	user, err := BizFindUserById(1)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("use not exists %+v\n", err)
			return
		}

		log.Printf("query user detail failed: %+v\n", err)
		return
	}

	log.Printf("user info: %+v\n", user)

}
