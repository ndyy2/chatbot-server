// config/database.go
package config

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMariaDB(uri string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(uri), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func InitMongoDB(uri string) (*mongo.Client, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	return client, nil
}