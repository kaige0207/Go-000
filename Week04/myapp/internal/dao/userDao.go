package dao

import (
	"database/sql"
	"fmt"
	"github.com/kaige0207/Go-000/Week04/myapp/internal/data"
	"github.com/kaige0207/Go-000/Week04/myapp/internal/pkg/errortype"
	"github.com/kaige0207/Go-000/Week04/myapp/internal/pkg/mysqldb"
	"github.com/pkg/errors"
	"log"
)

type UserDao struct {

}


func (dao *UserDao) GetUserByName(username string) (user *data.User, err error) {
	db, err := mysqldb.NewDB()
	if err != nil {
		log.Println("failed to open database:", err.Error())
		return nil,errors.Wrap(err, fmt.Sprintf("dao error: failed to open database: %+v", err.Error()))
	}
	defer db.Close()

	var id uint
	var password string
	row := db.QueryRow("SELECT age FROM users WHERE username = ?", username)
	if err = row.Scan(&id, &username, &password); err != nil {
		if err == sql.ErrNoRows {
			err = errortype.New(404, "data error: this user is not exist!")
		}
		return nil,errors.Wrap(err, fmt.Sprintf("dao error: find user by id=%+v", username))
	}
	user = &data.User{Id: id, Username: username, Password: password}
	return user, nil
}

func (dao *UserDao) AddUser(user *data.User) error {
	db, err := mysqldb.NewDB()
	if err != nil {
		log.Println("failed to open database:", err.Error())
		return errors.Wrap(err, fmt.Sprintf("dao error: failed to open database: %+v", err.Error()))
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO user(username,password) VALUES(?, ?)",user.Username,user.Password)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("dao error: insert data failed: %+v", err.Error()))
	}

	return nil
}

