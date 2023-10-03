package web

type CategoryCreateRequest struct {
	// Required Fields
	CategoryCode string `json:"subject" validate:"required"`
	Name         string `json:"title" validate:"required"`
}
