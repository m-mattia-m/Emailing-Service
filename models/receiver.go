package models

type Receiver struct {
	Name string `json:"name"`
	ReceiverRequst
}

type ReceiverRequst struct {
	Keyword     string `json:"keyword"`
	Email       string `json:"email"`
	Subject     string `json:"subject"`
	Message     string `json:"message"`
	SSL         bool   `json:"ssl"`
	ContentType string `json:"contentType"`
}
