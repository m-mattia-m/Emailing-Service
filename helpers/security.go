package helpers

import (
	"crypto/rand"
	"emailing-service/models"
	"encoding/base64"
	"errors"
	"math/big"
)

func GenerateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_@!<>=?+*#%&/(){}[]"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}

func GenerateRandomStringURLSafe(n int) (string, error) {
	b, err := GenerateRandomBytes(n)
	return base64.URLEncoding.EncodeToString(b), err
}

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func CheckEmailAttribute(email models.Email) error {
	if email.FromName == "" {
		return errors.New("an attribute is empty")
	}
	if email.FromEmail == "" {
		return errors.New("an attribute is empty")
	}
	if email.ToEmail == "" {
		return errors.New("an attribute is empty")
	}
	if email.Subject == "" {
		return errors.New("an attribute is empty")
	}
	if email.Message == "" {
		return errors.New("an attribute is empty")
	}
	if email.ContentType == "" {
		return errors.New("an attribute is empty")
	}
	if email.Host == "" {
		return errors.New("an attribute is empty")
	}
	if email.Port == "" {
		return errors.New("an attribute is empty")
	}
	if email.Username == "" {
		return errors.New("an attribute is empty")
	}
	if email.Password == "" {
		return errors.New("an attribute is empty")
	}
	return nil
}

func CheckRegisterAttribute(register models.SenderRequest) error {
	if register.Name == "" {
		return errors.New("an attribute is empty")
	}
	if register.Keyword == "" {
		return errors.New("an attribute is empty")
	}
	if register.Email == "" {
		return errors.New("an attribute is empty")
	}
	if register.NotificationEmail == "" {
		return errors.New("an attribute is empty")
	}
	if register.SmtpUrl == "" {
		return errors.New("an attribute is empty")
	}
	if register.SmtpPort == "" {
		return errors.New("an attribute is empty")
	}
	if register.SmtpUsername == "" {
		return errors.New("an attribute is empty")
	}
	if register.SmtpPassword == "" {
		return errors.New("an attribute is empty")
	}
	return nil
}
