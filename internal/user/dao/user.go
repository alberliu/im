package dao

import (
	"database/sql"
	"fmt"
	"im/internal/user/model"
	"im/pkg/db"
	"im/pkg/gerrors"
	"im/pkg/util"
)

type userDao struct{}

var UserDao = new(userDao)

// Add 插入一条用户信息
func (*userDao) Add(user model.User) (int64, error) {
	result, err := db.DBCli.Exec("insert ignore into user(phone_number,nickname,sex,avatar_url,extra) values(?,?,?,?,?)",
		user.PhoneNumber, user.Nickname, user.Sex, user.AvatarUrl, user.Extra)
	if err != nil {
		return 0, gerrors.WrapError(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, gerrors.WrapError(err)
	}
	return id, nil
}

// Get 获取用户信息
func (*userDao) Get(userId int64) (*model.User, error) {
	row := db.DBCli.QueryRow("select phone_number,nickname,sex,avatar_url,extra,create_time,update_time from user where id = ?",
		userId)
	user := model.User{
		Id: userId,
	}

	err := row.Scan(&user.PhoneNumber, &user.Nickname, &user.Sex, &user.AvatarUrl, &user.Extra, &user.CreateTime, &user.UpdateTime)
	if err != nil && err != sql.ErrNoRows {
		return nil, gerrors.WrapError(err)
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &user, err
}

// Get 获取用户信息
func (*userDao) GetByIds(userIds []int64) ([]model.User, error) {
	sql := fmt.Sprintf("select id,phone_number,nickname,sex,avatar_url,extra,create_time,update_time from user where id in %s", util.In(userIds))
	rows, err := db.DBCli.Query(sql)
	if err != nil {
		return nil, err
	}

	users := make([]model.User, 0, len(userIds))
	for rows.Next() {
		user := model.User{}
		err := rows.Scan(&user.Id, &user.PhoneNumber, &user.Nickname, &user.Sex, &user.AvatarUrl, &user.Extra, &user.CreateTime, &user.UpdateTime)
		if err != nil {
			return nil, gerrors.WrapError(err)
		}
		users = append(users, user)
	}
	return users, err
}

// GetByPhoneNumber 根据手机号获取用户信息
func (*userDao) GetByPhoneNumber(phomeNumber string) (*model.User, error) {
	row := db.DBCli.QueryRow("select id,nickname,sex,avatar_url,extra,create_time,update_time from user where phone_number = ?",
		phomeNumber)
	user := model.User{
		PhoneNumber: phomeNumber,
	}

	err := row.Scan(&user.Id, &user.Nickname, &user.Sex, &user.AvatarUrl, &user.Extra, &user.CreateTime, &user.UpdateTime)
	if err != nil && err != sql.ErrNoRows {
		return nil, gerrors.WrapError(err)
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &user, err
}

// Update 更新用户信息
func (*userDao) Update(user model.User) error {
	_, err := db.DBCli.Exec("update user set nickname = ?,sex = ?,avatar_url = ?,extra = ? where id = ?",
		user.Nickname, user.Sex, user.AvatarUrl, user.Extra, user.Id)
	if err != nil {
		return gerrors.WrapError(err)
	}

	return nil
}
