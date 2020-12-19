package dao

import (
	"fmt"
	"github.com/kaige0207/Go-000/Week04/myapp/internal/data"
	"testing"
)

func TestGetUserByName(t *testing.T) {
	fmt.Println(GetUserByName("nick"))
	fmt.Println(GetUserByName("jacky"))
	fmt.Println(GetUserByName("xxx"))
	fmt.Println(GetUserByName("aaa"))
}


func TestAddUser(t *testing.T) {
	user := &data.User{Username:"aaa",Password:"aaa"}
	fmt.Println(AddUser(user))
}
