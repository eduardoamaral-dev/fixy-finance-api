package models

import (
	"time"
)

type NewExpenseRequest struct {
	Title     string    `bson:"title,omitempty"`
	Category  string    `bson:"category,omitempty"`
	Price     float32   `bson:"price,omitempty"`
	CreatedAt time.Time `bson:"created_at,omitempty"`
}
