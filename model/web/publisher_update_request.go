package web

type PublisherUpdateRequest struct {
	// Required Fields
	PublisherCode string `json:"subject" validate:"required"`
	Name          string `json:"title" validate:"required"`
}
