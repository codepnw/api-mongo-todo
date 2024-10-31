package models

import "time"

type Todo struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
	Task      string    `json:"task,omitempty" bson:"task,omitempty"`
	Completed bool      `json:"completed,omitempty" bson:"completed,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
