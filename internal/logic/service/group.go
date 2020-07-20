package service

import (
	"context"
	"im/internal/logic/cache"
	"im/internal/logic/dao"
	"im/internal/logic/model"
	"im/pkg/gerrors"
	"im/pkg/pb"
	"im/pkg/rpc"
)

type groupService struct{}

var GroupService = new(groupService)

// Get 获取群组信息
func (*groupService) Get(ctx context.Context, groupId int64) (*model.Group, error) {
	group, err := cache.GroupCache.Get(groupId)
	if err != nil {
		return nil, err
	}
	if group != nil {
		return group, nil
	}
	group, err = dao.GroupDao.Get(groupId)
	if err != nil {
		return nil, err
	}
	err = cache.GroupCache.Set(group)
	if err != nil {
		return nil, err
	}
	return group, nil
}

// Create 创建群组
func (*groupService) Create(ctx context.Context, group model.Group) (int64, error) {
	return dao.GroupDao.Add(group)
}

// Update 更新群组
func (*groupService) Update(ctx context.Context, group model.Group) error {
	err := dao.GroupDao.Update(group.Id, group.Name, group.Introduction, group.Extra)
	if err != nil {
		return err
	}
	err = cache.GroupCache.Del(group.Id)
	if err != nil {
		return err
	}
	return nil
}

// AddUser 给群组添加用户
func (*groupService) AddUser(ctx context.Context, groupId, userId int64, remarks, extra string) error {
	group, err := GroupService.Get(ctx, groupId)
	if err != nil {
		return err
	}
	if group == nil {
		return gerrors.ErrGroupNotExist
	}

	if group.Type == model.GroupTypeGroup {
		err = GroupUserService.AddUser(ctx, groupId, userId, remarks, extra)
		if err != nil {
			return err
		}
	}
	if group.Type == model.GroupTypeChatRoom {
		err = cache.LargeGroupUserCache.Set(groupId, userId, remarks, extra)
		if err != nil {
			return err
		}
	}
	return nil
}

// UpdateUser 更新群组用户
func (*groupService) UpdateUser(ctx context.Context, groupId, userId int64, label, extra string) error {
	group, err := GroupService.Get(ctx, groupId)
	if err != nil {
		return err
	}

	if group == nil {
		return gerrors.ErrGroupNotExist
	}

	if group.Type == model.GroupTypeGroup {
		err = GroupUserService.Update(ctx, groupId, userId, label, extra)
		if err != nil {
			return err
		}
	}
	if group.Type == model.GroupTypeChatRoom {
		err = cache.LargeGroupUserCache.Set(groupId, userId, label, extra)
		if err != nil {
			return err
		}
	}
	return nil
}

// DeleteUser 删除用户群组
func (*groupService) DeleteUser(ctx context.Context, groupId, userId int64) error {
	group, err := GroupService.Get(ctx, groupId)
	if err != nil {
		return err
	}

	if group == nil {
		return gerrors.ErrGroupNotExist
	}

	if group.Type == model.GroupTypeGroup {
		err = GroupUserService.DeleteUser(ctx, groupId, userId)
		if err != nil {
			return err
		}
	}
	if group.Type == model.GroupTypeChatRoom {
		err = cache.LargeGroupUserCache.Del(groupId, userId)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetUsers 获取群组用户
func (s *groupService) GetUsers(ctx context.Context, groupId int64) ([]*pb.GroupMember, error) {
	group, err := s.Get(ctx, groupId)
	if err != nil {
		return nil, err
	}
	if group == nil {
		return nil, nil
	}
	if group.Type != model.GroupTypeGroup {
		return nil, nil
	}

	members, err := GroupUserService.GetUsers(ctx, groupId)
	if err != nil {
		return nil, err
	}

	userIds := make([]int64, len(members))
	for i := range members {
		userIds[i] = members[i].UserId
	}
	resp, err := rpc.UserIntClient.GetUsers(ctx, &pb.GetUsersReq{UserIds: userIds})
	if err != nil {
		return nil, err
	}

	var infos = make([]*pb.GroupMember, len(members))
	for i := range members {
		member := pb.GroupMember{
			UserId:  members[i].UserId,
			Remarks: members[i].Remarks,
			Extra:   members[i].Extra,
		}

		user, ok := resp.Users[members[i].UserId]
		if ok {
			member.Nickname = user.Nickname
			member.Sex = user.Sex
			member.AvatarUrl = user.AvatarUrl
			member.UserExtra = user.Extra
		}
		infos[i] = &member
	}

	return infos, nil
}
