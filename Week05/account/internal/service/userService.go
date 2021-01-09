package service

import (
	"context"
	pb "github.com/kaige0207/Go-000/Week05/account/api/v1/user"
	"github.com/kaige0207/Go-000/Week05/account/internal/dao"
	"github.com/kaige0207/Go-000/Week05/account/internal/data"
	"github.com/kaige0207/Go-000/Week05/account/internal/pkg/errortype"
	"github.com/pkg/errors"
	"log"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	UserDao *dao.UserDao
}

func (s *UserService) Login(ctx context.Context, in *pb.UserRequest) (*pb.UserReply, error) {

	log.Printf("Received: %v,%v", in.GetUsername(), in.GetPassword())
	user, err := s.UserDao.GetUserByName(in.GetUsername())
	var e = errortype.New(404, "")
	if err != nil {
		if errors.As(err, &e) {
			return &pb.UserReply{Message: "该用户未注册！"}, nil
		}
		log.Printf("query user failed: %+v\n", err)
	}

	if in.GetPassword() != user.Password {
		return &pb.UserReply{Message: "用户密码不正确！"}, nil
	}

	return &pb.UserReply{Message: "登录成功"}, nil
}

func (s *UserService) Register(ctx context.Context, in *pb.UserRequest) (*pb.UserReply, error) {
	user, err := s.UserDao.GetUserByName(in.GetUsername())
	if err != nil {
		log.Printf("query user failed: %+v\n", err)
	}

	if user != nil {
		return &pb.UserReply{Message: "该用户已存在！"}, nil
	}
	if err := s.UserDao.AddUser(&data.User{Username: in.GetUsername(), Password: in.GetPassword()}); err != nil {
		log.Printf("register user failed: %+v\n", err)
	}
	return &pb.UserReply{Message: "注册成功！"}, nil
}
