package db

import (
	"github.com/sakasikai/GoFive/pkg/constants"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Init init DB
func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(constants.MySQLDefaultDSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}

	// 手动建表
	if !DB.Migrator().HasTable(constants.UserTableName) {
		DB.Migrator().CreateTable(&User{})
	}

	// todo add gormopentracing
}
