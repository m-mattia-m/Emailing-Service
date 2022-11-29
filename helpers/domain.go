package helpers

import "strings"

func GetDomain(email string) string {
	emailParts := strings.Split(email, "@")
	if len(emailParts) == 0 {
		return ""
	}
	return emailParts[1]
}
