package dao

import (
	"database/sql"
	"im/internal/logic/model"
	"im/pkg/db"
	"im/pkg/gerrors"
	"im/pkg/logger"
)

type deviceDao struct{}

var DeviceDao = new(deviceDao)

// Insert 插入一条设备信息
func (*deviceDao) Add(device model.Device) (int64, error) {
	result, err := db.DBCli.Exec(`insert into device(type,brand,model,system_version,sdk_version,status,conn_addr) 
		values(?,?,?,?,?,?,?)`,
		device.Type, device.Brand, device.Model, device.SystemVersion, device.SDKVersion, device.Status, "")
	if err != nil {
		return 0, gerrors.WrapError(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, gerrors.WrapError(err)
	}
	return id, nil
}

// Get 获取设备
func (*deviceDao) Get(deviceId int64) (*model.Device, error) {
	device := model.Device{
		Id: deviceId,
	}
	row := db.DBCli.QueryRow(`
		select user_id,type,brand,model,system_version,sdk_version,status,conn_addr,create_time,update_time
		from device where id = ?`, deviceId)
	err := row.Scan(&device.UserId, &device.Type, &device.Brand, &device.Model, &device.SystemVersion, &device.SDKVersion,
		&device.Status, &device.ConnAddr, &device.CreateTime, &device.UpdateTime)
	if err != nil && err != sql.ErrNoRows {
		return nil, gerrors.WrapError(err)
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &device, err
}

// ListUserOnline 查询用户所有的在线设备
func (*deviceDao) ListOnlineByUserId(userId int64) ([]model.Device, error) {
	rows, err := db.DBCli.Query(
		`select id,type,brand,model,system_version,sdk_version,status,conn_addr,conn_fd,create_time,update_time from device where user_id = ? and status = ?`,
		userId, model.DeviceOnLine)
	if err != nil {
		return nil, gerrors.WrapError(err)
	}

	devices := make([]model.Device, 0, 5)
	for rows.Next() {
		device := new(model.Device)
		err = rows.Scan(&device.Id, &device.Type, &device.Brand, &device.Model, &device.SystemVersion, &device.SDKVersion,
			&device.Status, &device.ConnAddr, &device.ConnFd, &device.CreateTime, &device.UpdateTime)
		if err != nil {
			logger.Sugar.Error(err)
			return nil, err
		}
		devices = append(devices, *device)
	}
	return devices, nil
}

// UpdateUserIdAndStatus 更新设备绑定用户和设备在线状态
func (*deviceDao) UpdateUserIdAndStatus(deviceId, userId int64, status int, connectAddr string) error {
	_, err := db.DBCli.Exec("update device  set user_id = ?,status = ?,conn_addr = ? where id = ? ",
		userId, status, connectAddr, deviceId)
	if err != nil {
		return gerrors.WrapError(err)
	}
	return nil
}

// UpdateStatus 更新设备的在线状态
func (*deviceDao) UpdateStatus(deviceId int64, status int) error {
	_, err := db.DBCli.Exec("update device set status = ? where id = ?", status, deviceId)
	if err != nil {
		return gerrors.WrapError(err)
	}
	return nil
}

// Upgrade 升级设备
func (*deviceDao) Upgrade(deviceId int64, systemVersion, sdkVersion string) error {
	_, err := db.DBCli.Exec("update device set system_version = ?,sdk_version = ? where id = ? ",
		systemVersion, sdkVersion, deviceId)
	if err != nil {
		return gerrors.WrapError(err)
	}
	return nil
}
