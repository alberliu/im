package dao

import (
	"database/sql"
	"im/internal/logic/model"
	"im/pkg/db"
	"im/pkg/gerrors"
)

type groupDao struct{}

var GroupDao = new(groupDao)

// Get 获取群组信息
func (*groupDao) Get(groupId int64) (*model.Group, error) {
	row := db.DBCli.QueryRow("select name,introduction,user_num,type,extra,create_time,update_time from `group` where id = ?",
		groupId)
	group := model.Group{
		Id: groupId,
	}
	err := row.Scan(&group.Name, &group.Introduction, &group.UserNum, &group.Type, &group.Extra, &group.CreateTime, &group.UpdateTime)
	if err != nil && err != sql.ErrNoRows {
		return nil, gerrors.WrapError(err)
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &group, nil
}

// Insert 插入一条群组
func (*groupDao) Add(group model.Group) (int64, error) {
	result, err := db.DBCli.Exec("insert ignore into `group`(name,introduction,type,extra) value(?,?,?,?)",
		group.Name, group.Introduction, group.Type, group.Extra)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, gerrors.WrapError(err)
	}
	return id, nil
}

// Update 更新群组信息
func (*groupDao) Update(groupId int64, name, introduction, extra string) error {
	_, err := db.DBCli.Exec("update `group` set name = ?,introduction = ?,extra = ? where id = ?",
		name, introduction, extra, groupId)
	if err != nil {
		return gerrors.WrapError(err)
	}

	return nil
}

// UpdateUserNum 更新群组信息
func (*groupDao) UpdateUserNum(groupId int64, userNum int) error {
	_, err := db.DBCli.Exec("update `group` set user_num = user_num + ? where id = ?",
		userNum, groupId)
	if err != nil {
		return gerrors.WrapError(err)
	}

	return nil
}
