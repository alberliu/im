package api

import (
	"context"
	"im/internal/user/service"
	"im/pkg/pb"
)

type UserIntServer struct{}

func (*UserIntServer) Auth(ctx context.Context, req *pb.AuthReq) (*pb.AuthResp, error) {
	return &pb.AuthResp{}, service.AuthService.Auth(ctx, req.UserId, req.DeviceId, req.Token)
}
func (*UserIntServer) GetUsers(ctx context.Context, req *pb.GetUsersReq) (*pb.GetUsersResp, error) {
	users, err := service.UserService.GetByIds(ctx, req.UserIds)
	if err != nil {
		return nil, err
	}

	pbUsers := make([]*pb.User, 0, len(users))
	for i := range users {
		pbUsers = append(pbUsers, &pb.User{
			UserId:     users[i].Id,
			Nickname:   users[i].Nickname,
			Sex:        users[i].Sex,
			AvatarUrl:  users[i].AvatarUrl,
			Extra:      users[i].Extra,
			CreateTime: users[i].CreateTime.Unix(),
			UpdateTime: users[i].UpdateTime.Unix(),
		})
	}

	return &pb.GetUsersResp{Users: pbUsers}, nil
}
