package models

import "gorm.io/gorm"

type DummyUsers struct {
	ID   uint   `gorm:"primary key;autoIncrement" json:"id"`
	Hash string `json:"hash"`
}

func MigrateDummyUsers(db *gorm.DB) error {
	err := db.AutoMigrate(&DummyUsers{})
	return err
}
