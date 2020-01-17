package service

import (
	"context"
	"im/internal/logic/cache"
	"im/internal/logic/dao"
	"im/internal/logic/model"
	"im/pkg/gerrors"
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
func (*groupService) Create(ctx context.Context, group model.Group) error {
	affected, err := dao.GroupDao.Add(group)
	if err != nil {
		return err
	}

	if affected == 0 {
		return gerrors.ErrGroupAlreadyExist
	}
	return nil
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
func (*groupService) AddUser(ctx context.Context, groupId, userId int64, label, extra string) error {
	group, err := GroupService.Get(ctx, groupId)
	if err != nil {
		return err
	}
	if group == nil {
		return gerrors.ErrGroupNotExist
	}

	if group.Type == model.GroupTypeGroup {
		err = GroupUserService.AddUser(ctx, groupId, userId, label, extra)
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
