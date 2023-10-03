package web

type BookUpdateRequest struct {
	// Required Fields
	BookCode string `json:"subject" validate:"required"`
	Title    string `json:"title" validate:"required"`
}
