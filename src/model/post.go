package model

// Post model
// `json:"...."` StrucTag -> https://golang.org/pkg/reflect/#StructTag
type Post struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Text     string `json:"text"`
	Category string `json:"category"`
	Message  string `json:"message"`
}
