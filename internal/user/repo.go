package user

import "go.mongodb.org/mongo-driver/mongo"

type Repo struct {
	col *mongo.Collection
}


func NewRepo(db *mongo.Database) *Repo {
	return &Repo{
		col: db.Collection("users"),
	}
}