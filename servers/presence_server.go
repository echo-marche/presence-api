package servers

import (
	"context"

	"github.com/alduh-co/presence-api/models"
	pb "github.com/alduh-co/presence-api/proto/pb"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PresenceServer struct {
	Db *gorm.DB
}

func (server *PresenceServer) GetUserList(ctx context.Context, req *pb.UserListRequest) (*pb.UserListResponse, error) {
	responseUsers := []*pb.UserResponse{}
	var users models.Users
	if err := server.Db.Find(&users).Error; err != nil {
		return nil, status.Errorf(codes.DataLoss,
			err.Error())
	}
	for _, user := range users {
		responseUsers = append(responseUsers, &pb.UserResponse{Name: user.Name, Email: user.Email})
	}
	return &pb.UserListResponse{Users: responseUsers}, nil
}

func (server *PresenceServer) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	// systemCode := req.GetSystemCode
	return &pb.UserResponse{Name: "test1", Email: "test1@gmail.com"}, nil
}
