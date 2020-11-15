package db

import (

	"fmt"

	"time"
	"github.com/jinzhu/gorm"
)

var Driver driver

func InitDB() {

Conn:
	o, err := gorm.Open("mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=Local",
			"root",
			"12345678",
			"mariadb",
			"3306",
			"mysql",
		))
	Driver.DB = o
	if err != nil {
		fmt.Println("[!] 数据库链接异常，尝试重新链接...", err)
		time.Sleep(time.Second * 5)
		goto Conn
	}
	Driver.Callback().Create().Before("gorm:create").Register("update_time_in_create", updateTimeForCreateCallback)
	Driver.Callback().Update().Before("gorm:update").Register("update_time_in_update", updateTimeForUpdateCallback)

	Driver.DB.DB().SetMaxIdleConns(30)
	Driver.DB.DB().SetMaxOpenConns(500)
	Driver.SingularTable(true)

	Driver.AutoMigrate(
		&Test{},

	)
}

// 更新创建时间
func updateTimeForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		if scope.HasColumn("CreateTime") {
			scope.SetColumn("CreateTime", time.Now().Unix())
		}
		updateTimeForUpdateCallback(scope)
	}
}

// 更新更新时间
func updateTimeForUpdateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		// 更新时间
		if scope.HasColumn("UpdateTime") {
			scope.SetColumn("UpdateTime", time.Now().Unix())
		}
	}
}


type driver struct {
	*gorm.DB
}

