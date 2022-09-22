package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func init() {
	dbConnect := sqlite.Open("db/db.db")

	var err error
	DB, err = gorm.Open(dbConnect, &gorm.Config{})
	if err != nil {
		log.Fatal("数据库连接失败...")
	}

	migrate(DB)
}

func migrate(db *gorm.DB) {
	db.Set("gorm:table_options", "AUTO_INCREMENT=100")
	db.AutoMigrate(&User{})

	db.AutoMigrate(&Album{})
	db.AutoMigrate(&Artist{})
	db.AutoMigrate(&Comment{})
	db.AutoMigrate(&Playlist{})
	db.AutoMigrate(&PlaylistItems{})
	db.AutoMigrate(&Song{})
}
