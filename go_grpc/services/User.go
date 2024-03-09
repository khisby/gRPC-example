package services

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/khisby/go-grpc-example/grpc/User"
)

type User struct {
	UserId   string
	Username string
}

type UserService struct {
	pb.UnimplementedUserServiceServer
	users []*User
}

func NewUser() *UserService {

	users := []*User{}
	users = append(users, &User{UserId: "khisoft", Username: "Khisby"})

	return &UserService{
		UnimplementedUserServiceServer: pb.UnimplementedUserServiceServer{},
		users:                          users,
	}
}

func (s *UserService) CreateUser(ctx context.Context, in *pb.User) (*pb.UserResponse, error) {
	if in.UserId == "" {
		return &pb.UserResponse{
			Message: "UserId is required",
			Status:  "error",
		}, nil
	}

	if in.Username == "" {
		return &pb.UserResponse{
			Message: "Username is required",
			Status:  "error",
		}, nil
	}

	s.users = append(s.users, &User{
		UserId:   in.UserId,
		Username: in.Username,
	})

	return &pb.UserResponse{
		Data:    in,
		Message: "Success create user",
		Status:  "success",
	}, nil
}

func (s *UserService) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.UserResponse, error) {
	var user *User
	for _, u := range s.users {
		if u.UserId == in.UserId {
			user = u
			break
		}
	}

	if user == nil {
		return &pb.UserResponse{
			Message: "User not found",
			Status:  "error",
		}, nil
	}

	return &pb.UserResponse{
		Data:    parseToUserResponse(user),
		Message: "Success find user by UserId",
		Status:  "success",
	}, nil
}

func (s *UserService) GetUsers(ctx context.Context, in *empty.Empty) (*pb.UsersResponse, error) {
	return &pb.UsersResponse{
		Data:    parseToUsersResponse(s.users),
		Message: "Success to get all users",
		Status:  "sucess",
	}, nil
}

func parseToUsersResponse(users []*User) []*pb.User {
	var userResponses []*pb.User
	for _, user := range users {
		userResponses = append(userResponses, parseToUserResponse(user))
	}
	return userResponses
}

func parseToUserResponse(user *User) *pb.User {
	return &pb.User{
		UserId:   user.UserId,
		Username: user.Username,
	}
}
