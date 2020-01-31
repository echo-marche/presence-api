package servers

import (
	"context"

	pb "github.com/alduh-co/presence-api/proto/pb"
)

type UserServer struct{}

func (s *UserServer) GetUserList(ctx context.Context, req *pb.UserListRequest) (*pb.UserListResponse, error) {
	users := []*pb.UserResponse{}
	users = append(users, &pb.UserResponse{Name: "test1", Email: "test1@gmail.com"})
	users = append(users, &pb.UserResponse{Name: "test2", Email: "test2@gmail.com"})
	users = append(users, &pb.UserResponse{Name: "test3", Email: "test3@gmail.com"})
	return &pb.UserListResponse{Users: users}, nil
}

func (s *UserServer) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	// systemCode := req.GetSystemCode
	return &pb.UserResponse{Name: "test1", Email: "test1@gmail.com"}, nil
}
