package models

import "gorm.io/gorm"

func SetupTriggers(conn *gorm.DB) {
	SetupSearchIndexTrigger(conn)
}

func GetAll() []interface{} {
	return []interface{}{
		&User{}, &Card{}, &Folder{},
	}
}
