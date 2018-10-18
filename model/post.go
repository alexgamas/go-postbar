package model

// Post model
// `json:"...."` StrucTag -> https://golang.org/pkg/reflect/#StructTag
type Post struct {
	ID       string `json:"_id,omitempty"`
	Title    string `json:"title,omitempty"`
	Text     string `json:"text,omitempty"`
	Category string `json:"category,omitempty"`
	Message  string `json:"message,omitempty"`
}
