package web

type BookCreateRequest struct {
	// Required Fields
	BookCode string `json:"subject" validate:"required"`
	Title    string `json:"title" validate:"required"`
}
