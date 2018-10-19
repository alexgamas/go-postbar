package model

import (
	"time"
)

// Post model
// `json:"...."` StrucTag -> https://golang.org/pkg/reflect/#StructTag
type Post struct {
	ID      string    `json:"_id,omitempty"`
	Title   string    `json:"title,omitempty"`
	Message string    `json:"message,omitempty"`
	Created time.Time `json:"created,omitempty"`
}
