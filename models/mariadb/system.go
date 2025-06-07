package mariadb

import "gorm.io/gorm"

type SystemSetting struct {
	gorm.Model
	Key   string `gorm:"uniqueIndex"`
	Value string
}

type AuditLog struct {
	gorm.Model
	Action    string
	UserID    uint
	IPAddress string
}