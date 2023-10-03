package web

type CategoryUpdateRequest struct {
	// Required Fields
	CategoryCode string `json:"subject" validate:"required"`
	Name         string `json:"title" validate:"required"`
}
