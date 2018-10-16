package model

// Health Check model
// `json:"...."` StrucTag -> https://golang.org/pkg/reflect/#StructTag
type Health struct {
	Status string `json:"status"`
}
