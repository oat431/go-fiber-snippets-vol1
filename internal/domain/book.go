package domain

// Book is the core domain model — framework-agnostic.
type Book struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}
