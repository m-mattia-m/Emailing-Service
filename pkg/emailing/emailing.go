package emailing

import (
	"crypto/tls"
	"emailing-service/helpers"
	"emailing-service/models"
	"fmt"
	"log"
	"net"
	"net/mail"
	"net/smtp"
)

func Mail(sender models.Sender, receiver models.ReceiverRequst) (models.Email, error) {
	email := models.Email{
		FromName:    sender.Name,
		FromEmail:   sender.Email,
		ToEmail:     receiver.Email,
		Subject:     receiver.Subject,
		Message:     receiver.Message,
		ContentType: receiver.ContentType,
		Host:        sender.SmtpUrl,
		Port:        sender.SmtpPort,
		Username:    sender.SmtpUsername,
		Password:    sender.SmtpPassword,
		SSL:         receiver.SSL,
	}

	notificationEmail := email
	notificationEmail.ToEmail = sender.NotificationEmail

	err := helpers.CheckEmailAttribute(email)
	if err != nil {
		return models.Email{}, err
	}
	err = helpers.CheckEmailAttribute(notificationEmail)
	if err != nil {
		return models.Email{}, err
	}

	sendMail(email)
	sendMail(notificationEmail)
	return email, nil
}

func sendMail(email models.Email) {
	from := mail.Address{Name: email.FromName, Address: email.FromEmail}
	to := mail.Address{Name: "", Address: email.ToEmail}
	subject := email.Subject
	body := email.Message

	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subject
	headers["Content-Type"] = email.ContentType

	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	servername := email.Host + ":" + email.Port
	host, _, _ := net.SplitHostPort(servername)
	auth := smtp.PlainAuth("", email.Username, email.Password, email.Host)

	tlsconfig := &tls.Config{
		InsecureSkipVerify: email.SSL,
		ServerName:         host,
	}

	conn, err := tls.Dial("tcp", servername, tlsconfig)
	if err != nil {
		log.Println("[MAIL]: Cannot connect to the mail server \t\t -> Error: " + err.Error())
		return
	}

	c, err := smtp.NewClient(conn, host)
	if err != nil {
		log.Println("[MAIL]: No mail client can be created \t\t -> Error: " + err.Error())
		return
	}

	// Auth
	if err = c.Auth(auth); err != nil {
		log.Println("[MAIL]: Authentication failed \t\t -> Error: " + err.Error())
		return
	}

	// To && From
	if err = c.Mail(from.Address); err != nil {
		log.Println("[MAIL]: Create Mail Failed \t\t -> Error: " + err.Error())
		return
	}

	if err = c.Rcpt(to.Address); err != nil {
		log.Println("[MAIL]: RCPT request failed \t\t -> Error: " + err.Error())
		return
	}

	// Data
	w, err := c.Data()
	if err != nil {
		log.Println("[MAIL]: Writing content to the mail failed \t\t -> Error: " + err.Error())
		return
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		log.Println("[MAIL]: Body writing in the mail failed \t\t -> Error: " + err.Error())
		return
	}

	err = w.Close()
	if err != nil {
		log.Println("[MAIL]: Close mail failed \t\t -> Error: " + err.Error())
		return
	}

	c.Quit()
}
