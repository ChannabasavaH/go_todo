package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

dsn := "host:localhost user=postgres password=Cha010@#3 dbname=todo port=9920 sslmode=disable TimeZone=Asia/Shanghai"
db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})