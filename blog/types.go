package main

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Blog struct {
	Id        bson.ObjectId `bson:"_id"`
	Title     string        `bson:"title"`
	Content   string        `bson:"content"`
	CreatedAt time.Time     `bson:"created_at"`
}
