package dao

import (
	"fmt"
	"github.com/kaige0207/Go-000/Week04/account/internal/data"
	"testing"
)

func TestGetUserByName(t *testing.T) {
	dao := &UserDao{}
	fmt.Println(dao.GetUserByName("nick"))
	fmt.Println(dao.GetUserByName("jacky"))
	fmt.Println(dao.GetUserByName("xxx"))
	fmt.Println(dao.GetUserByName("kaige"))
}

func TestAddUser(t *testing.T) {
	dao := &UserDao{}
	user := &data.User{Username: "aaa", Password: "aaa"}
	fmt.Println(dao.AddUser(user))
}
