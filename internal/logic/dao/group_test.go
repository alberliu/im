package dao

import (
	"fmt"
	"im/internal/logic/model"
	"testing"
)

func TestGroupDao_Get(t *testing.T) {
	fmt.Println(GroupDao.Get(1))
}

func TestGroupDao_Add(t *testing.T) {
	group := model.Group{
		Name:         "5",
		Introduction: "5",
		Type:         5,
		Extra:        "5",
	}
	fmt.Println(GroupDao.Add(group))
}

func TestGroupDao_Update(t *testing.T) {
	fmt.Println(GroupDao.Update(4, "4", "4", "4"))
}

func TestGroupDao_AddUserNum(t *testing.T) {
	fmt.Println(GroupDao.AddUserNum(4, 1))
}
