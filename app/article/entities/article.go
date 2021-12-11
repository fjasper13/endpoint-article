package entities

import "time"

type Article struct {
	ID        int64      `db:"id" json:"id"`
	Author    string     `db:"author" json:"author"`
	Title     string     `db:"title" json:"title"`
	Body      string     `db:"body" json:"body"`
	CreatedAt time.Time  `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
}
