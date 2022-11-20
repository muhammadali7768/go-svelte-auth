package dtos

type Response struct {
	Status  int64  `json:"status,omitempty"`
	Message string `json:"messge,omitempty"`
}
