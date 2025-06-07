// repositories/mongodb/session_repo.go
package mongodb

import (
	"ai-assistant/models/mongodb"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type SessionRepository struct {
	collection *mongo.Collection
}

func NewSessionRepository(db *mongo.Database) *SessionRepository {
	return &SessionRepository{
		collection: db.Collection("sessions"),
	}
}

func (r *SessionRepository) CreateSession(session *mongodb.Session) error {
	_, err := r.collection.InsertOne(context.Background(), session)
	return err
}