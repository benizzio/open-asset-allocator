package model

type ErrorResponse struct {
	ErrorMessage string   `json:"errorMessage"`
	Details      []string `json:"details,omitempty"`
}
