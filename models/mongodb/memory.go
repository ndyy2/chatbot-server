// models/mongodb/memory.go
package mongodb

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Memory struct {
	ID        primitive.ObjectID `bson:"_id"`
	SessionID string             `bson:"session_id"`
	Content   string             `bson:"content"`
	Vector    []float32          `bson:"vector"`
	CreatedAt time.Time          `bson:"created_at"`
}