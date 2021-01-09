package dao

import (
	"database/sql"
	"fmt"
	"github.com/kaige0207/Go-000/Week05/account/internal/data"
	"github.com/kaige0207/Go-000/Week05/account/internal/pkg/errortype"
	"github.com/kaige0207/Go-000/Week05/account/internal/pkg/mysqldb"
	"github.com/pkg/errors"
	"log"
	"time"
)

type UserDao struct {
	user data.User
}

func (dao *UserDao) GetUserByName(username string) (user *data.User, err error) {
	db, err := mysqldb.NewDB()
	if err != nil {
		log.Println("failed to open database:", err.Error())
		return nil, errors.Wrap(err, fmt.Sprintf("dao error: failed to open database: %+v", err.Error()))
	}

	var id uint
	var password string
	row := db.QueryRow("SELECT id,username,password FROM user WHERE username = ?", username)
	err = row.Scan(&id, &username, &password)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = errortype.New(404, "this user is not exit")
		}
		return nil, errors.Wrap(err, err.Error())
	}

	dao.user.Id = id
	dao.user.Username = username
	dao.user.Password = password
	return &dao.user, row.Err()
}

func (dao *UserDao) AddUser(user *data.User) error {
	db, err := mysqldb.NewDB()
	if err != nil {
		log.Println("failed to open database:", err.Error())
		return errors.Wrap(err, fmt.Sprintf("dao error: failed to open database: %+v", err.Error()))
	}

	_, err = db.Exec("INSERT INTO user(username,password,createtime,updtetime) VALUES(?, ?, ?, ?)", user.Username, user.Password, time.Now(), time.Now())
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("dao error: insert data failed: %+v", err.Error()))
	}

	return nil
}
