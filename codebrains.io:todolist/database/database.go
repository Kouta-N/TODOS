package database

import (
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	//構造体DBをDBConnに代入
	DBConn *gorm.DB
)
