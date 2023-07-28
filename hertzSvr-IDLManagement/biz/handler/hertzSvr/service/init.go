package service

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type IDLMessage struct {
	ID      int
	SvcName string
	IDL     string
}

var DB, _ = gorm.Open(sqlite.Open("IDLMessage.db"), &gorm.Config{})
