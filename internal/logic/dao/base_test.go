package dao

import (
	"fmt"
	"im/pkg/db"
)

func init() {
	fmt.Println("init db")
	db.InitByTest()
}
