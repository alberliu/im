package dao

import (
	"fmt"
	"im/internal/user/model"
	"im/pkg/util"
	"reflect"
	"testing"
)

func TestUserDao_Add(t *testing.T) {
	type args struct {
		user model.User
	}
	tests := []struct {
		name    string
		u       *userDao
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userDao{}
			got, err := u.Add(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("userDao.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("userDao.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserDao_Get(t *testing.T) {
	type args struct {
		userId int64
	}
	tests := []struct {
		name    string
		u       *userDao
		args    args
		want    *model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userDao{}
			got, err := u.Get(tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("userDao.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userDao.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserDao_GetByIds(t *testing.T) {
	fmt.Println(util.In([]int64{1, 2, 3}))
	fmt.Println(UserDao.GetByIds([]int64{1, 2, 3}))
}

func TestUserDao_GetByPhoneNumber(t *testing.T) {
	type args struct {
		phomeNumber string
	}
	tests := []struct {
		name    string
		u       *userDao
		args    args
		want    *model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userDao{}
			got, err := u.GetByPhoneNumber(tt.args.phomeNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("userDao.GetByPhoneNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userDao.GetByPhoneNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserDao_Update(t *testing.T) {
	type args struct {
		user model.User
	}
	tests := []struct {
		name    string
		u       *userDao
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userDao{}
			if err := u.Update(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("userDao.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
