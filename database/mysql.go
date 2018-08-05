package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql" //加载mysql
	"fmt"
)

var Eloquent *gorm.DB

type User struct {
	ID       int64  `json:"id"`       // 列名为 `id`
	Username string `json:"username"` // 列名为 `username`
	Password string `json:"password"` // 列名为 `password`
}

func init() {
	var err error
	Eloquent, err = gorm.Open("mysql", "homestead:secret@tcp(192.168.33.10:3306)/test?charset=utf8&parseTime=True&loc=Local&timeout=10ms")

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	if Eloquent.Error != nil {
		fmt.Printf("database error %v", Eloquent.Error)
	}

	Eloquent.AutoMigrate(&User{})

	// Add table suffix when create tables
	Eloquent.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})
}