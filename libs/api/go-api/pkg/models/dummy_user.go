package models

import "gorm.io/gorm"

type DummyUser struct {
	ID   uint        `gorm:"primary key;autoIncrement" json:"id"`
	Hash interface{} `json:"hash"`
}

func MigrateDummyUsers(db *gorm.DB) error {
	err := db.AutoMigrate(&DummyUser{})
	return err
}
