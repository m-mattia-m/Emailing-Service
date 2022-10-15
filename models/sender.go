package models

import "gorm.io/gorm"

type Sender struct {
	*gorm.Model `swaggerignore:"true"`
	SenderRequest
	Token string `json:"token"`
}

type SenderRequest struct {
	Name              string `json:"name"`
	Keyword           string `json:"keyword"`
	Email             string `json:"email"`
	NotificationEmail string `json:"notificationEmail"`
	SmtpUrl           string `json:"smtpUrl"`
	SmtpPort          string `json:"smtpPort"`
	SmtpUsername      string `json:"smtpUsername"`
	SmtpPassword      string `json:"smtpPassword"`
}
