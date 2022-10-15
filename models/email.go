package models

type Email struct {
	FromName    string
	FromEmail   string
	ToEmail     string
	Subject     string
	Message     string
	ContentType string
	Host        string
	Port        string
	Username    string
	Password    string
	SSL         bool
}
