// models/mariadb/user.go
package mariadb

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"uniqueIndex"`
	Password string
	Role     string // admin/user
}