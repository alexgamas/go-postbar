package model

import (
	"go-postbar/logger"
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

// Like Adiciona um like
func (p Post) Like() {
	logger.Log.Debug("%s", p)
}
