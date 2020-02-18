package model

// Request struct
type Request struct {
	Number  int    `json:"number" validate:"required"`
	Type    string `json:"type" validate:"required"`
	Version int    `json:"version"`
}
