package api

import (
	"context"
	"im/internal/logic/model"
	"im/internal/logic/service"
	"im/pkg/gerrors"
	"im/pkg/grpclib"
	"im/pkg/logger"
	"im/pkg/pb"
)

type LogicExtServer struct{}

// RegisterDevice 注册设备
func (*LogicExtServer) RegisterDevice(ctx context.Context, req *pb.RegisterDeviceReq) (*pb.RegisterDeviceResp, error) {
	device := model.Device{
		Type:          req.Type,
		Brand:         req.Brand,
		Model:         req.Model,
		SystemVersion: req.SystemVersion,
		SDKVersion:    req.SdkVersion,
	}

	if device.Type == 0 || device.Brand == "" || device.Model == "" ||
		device.SystemVersion == "" || device.SDKVersion == "" {
		return nil, gerrors.ErrBadRequest
	}

	id, err := service.DeviceService.Register(ctx, device)
	if err != nil {
		return nil, err
	}
	return &pb.RegisterDeviceResp{DeviceId: id}, nil
}

// SendMessage 发送消息
func (*LogicExtServer) SendMessage(ctx context.Context, req *pb.SendMessageReq) (*pb.SendMessageResp, error) {
	userId, deviceId, err := grpclib.GetCtxData(ctx)
	if err != nil {
		return nil, err
	}

	sender := model.Sender{
		SenderType: pb.SenderType_ST_USER,
		SenderId:   userId,
		DeviceId:   deviceId,
	}
	err = service.MessageService.Send(ctx, sender, *req)
	if err != nil {
		return nil, err
	}
	return &pb.SendMessageResp{}, nil
}

// CreateGroup 创建群组
func (*LogicExtServer) CreateGroup(ctx context.Context, req *pb.CreateGroupReq) (*pb.CreateGroupResp, error) {
	var group = model.Group{
		Name:         req.Group.Name,
		Introduction: req.Group.Introduction,
		Type:         req.Group.Type,
		Extra:        req.Group.Extra,
	}
	err := service.GroupService.Create(ctx, group)
	if err != nil {
		return nil, err
	}
	return &pb.CreateGroupResp{}, nil
}

// UpdateGroup 更新群组
func (*LogicExtServer) UpdateGroup(ctx context.Context, req *pb.UpdateGroupReq) (*pb.UpdateGroupResp, error) {

	var group = model.Group{
		Id:           req.Group.GroupId,
		Name:         req.Group.Name,
		Introduction: req.Group.Introduction,
		Type:         req.Group.Type,
		Extra:        req.Group.Extra,
	}
	err := service.GroupService.Update(ctx, group)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateGroupResp{}, nil
}

// GetGroup 获取群组信息
func (*LogicExtServer) GetGroup(ctx context.Context, req *pb.GetGroupReq) (*pb.GetGroupResp, error) {
	group, err := service.GroupService.Get(ctx, req.GroupId)
	if err != nil {
		return nil, err
	}

	if group == nil {
		return nil, gerrors.ErrGroupNotExist
	}

	return &pb.GetGroupResp{
		Group: &pb.Group{
			GroupId:      group.Id,
			Name:         group.Name,
			Introduction: group.Introduction,
			UserMum:      group.UserNum,
			Type:         group.Type,
			Extra:        group.Extra,
			CreateTime:   group.CreateTime.Unix(),
			UpdateTime:   group.UpdateTime.Unix(),
		},
	}, nil
}

// GetUserGroups 获取用户加入的所有群组
func (*LogicExtServer) GetUserGroups(ctx context.Context, in *pb.GetUserGroupsReq) (*pb.GetUserGroupsResp, error) {
	userId, _, err := grpclib.GetCtxData(ctx)
	if err != nil {
		return nil, err
	}

	groups, err := service.GroupUserService.ListByUserId(ctx, userId)
	if err != nil {
		logger.Sugar.Error(err)
		return nil, err
	}
	pbGroups := make([]*pb.Group, 0, len(groups))
	for i := range groups {
		pbGroups = append(pbGroups, &pb.Group{
			GroupId:      groups[i].Id,
			Name:         groups[i].Name,
			Introduction: groups[i].Introduction,
			UserMum:      groups[i].UserNum,
			Type:         groups[i].Type,
			Extra:        groups[i].Extra,
			CreateTime:   groups[i].CreateTime.Unix(),
			UpdateTime:   groups[i].UpdateTime.Unix(),
		})
	}
	return &pb.GetUserGroupsResp{Groups: pbGroups}, err
}

// AddGroupMember 添加群组成员
func (*LogicExtServer) AddGroupMember(ctx context.Context, in *pb.AddGroupMemberReq) (*pb.AddGroupMemberResp, error) {
	err := service.GroupService.AddUser(ctx, in.GroupUser.GroupId, in.GroupUser.UserId, in.GroupUser.Label, in.GroupUser.Extra)
	if err != nil {
		logger.Sugar.Error(err)
		return nil, err
	}

	return &pb.AddGroupMemberResp{}, nil
}

// UpdateGroupMember 更新群组成员信息
func (*LogicExtServer) UpdateGroupMember(ctx context.Context, in *pb.UpdateGroupMemberReq) (*pb.UpdateGroupMemberResp, error) {
	err := service.GroupService.UpdateUser(ctx, in.GroupUser.GroupId, in.GroupUser.UserId, in.GroupUser.Label, in.GroupUser.Extra)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateGroupMemberResp{}, nil
}

// DeleteGroupMember 添加群组成员
func (*LogicExtServer) DeleteGroupMember(ctx context.Context, in *pb.DeleteGroupMemberReq) (*pb.DeleteGroupMemberResp, error) {
	err := service.GroupService.DeleteUser(ctx, in.GroupId, in.UserId)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteGroupMemberResp{}, nil
}
