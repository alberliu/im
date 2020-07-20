package service

import (
	"context"
	"im/internal/logic/cache"
	"im/internal/logic/dao"
	"im/internal/logic/model"
)

type groupUserService struct{}

var GroupUserService = new(groupUserService)

// ListByUserId 获取用户所加入的群组
func (*groupUserService) ListByUserId(ctx context.Context, userId int64) ([]model.Group, error) {
	groups, err := dao.GroupUserDao.ListByUserId(userId)
	if err != nil {
		return nil, err
	}
	return groups, nil
}

// GetUsers 获取群组的所有用户信息
func (*groupUserService) GetUsers(ctx context.Context, groupId int64) ([]model.GroupUser, error) {
	users, err := cache.GroupUserCache.Get(groupId)
	if err != nil {
		return nil, err
	}

	if users != nil {
		return users, nil
	}

	users, err = dao.GroupUserDao.ListUser(groupId)
	if err != nil {
		return nil, err
	}

	err = cache.GroupUserCache.Set(groupId, users)
	if err != nil {
		return nil, err
	}
	return users, err
}

// AddUser 给群组添加用户
func (*groupUserService) AddUser(ctx context.Context, groupId, userId int64, remarks, extra string) error {
	err := dao.GroupUserDao.Add(groupId, userId, remarks, extra)
	if err != nil {
		return err
	}

	err = dao.GroupDao.UpdateUserNum(groupId, 1)
	if err != nil {
		return err
	}

	err = cache.GroupUserCache.Del(groupId)
	if err != nil {
		return err
	}

	return nil
}

// DeleteUser 从群组移除用户
func (*groupUserService) DeleteUser(ctx context.Context, groupId, userId int64) error {
	err := dao.GroupUserDao.Delete(groupId, userId)
	if err != nil {
		return err
	}

	err = dao.GroupDao.UpdateUserNum(groupId, -1)
	if err != nil {
		return err
	}

	err = cache.GroupUserCache.Del(groupId)
	if err != nil {
		return err
	}

	return nil
}

// Update 更新群组用户信息
func (*groupUserService) Update(ctx context.Context, groupId int64, userId int64, remarks, extra string) error {
	err := dao.GroupUserDao.Update(groupId, userId, remarks, extra)
	if err != nil {
		return err
	}

	err = cache.GroupUserCache.Del(groupId)
	if err != nil {
		return err
	}
	return nil
}
