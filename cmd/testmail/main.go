package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"mime"
	"net/smtp"
	"os"
	"strconv"

	"github.com/hay-kot/easyemails"
)

func EnvOrDefault(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

func main() {
	var (
		SMTPHost   = EnvOrDefault("SMTP_HOST", "localhost")
		SMTPSender = EnvOrDefault("SMTP_SENDER_EMAIL", "EasyMail Test")
	)

	mailer := &Mailer{
		Host: SMTPHost,
		Port: 1025,
		From: SMTPSender,
	}

	bldr := easyemails.NewBuilder().Add(
		easyemails.WithParagraph(
			easyemails.WithText("Hello, world!"),
			easyemails.WithLineBreak(),
			easyemails.WithText("This is a test email, it works with [markdown](http://example.com)."),
			easyemails.WithList(
				"[Google](http://google.com) is a search engine.",
				"Item 2",
				"Item 3",
			),
			easyemails.WithLineBreak(),
			easyemails.WithText("I supported **bold** and *italic* text."),
			easyemails.WithLineBreak(),
		),
		easyemails.WithButton("Click me", "http://example.com"),
		easyemails.WithParagraph(
			easyemails.WithText("[Website](http://example.com/website) · [Unsubscribe](http://example.com/unsubscribe)").Centered(),
		).FontSize(12),
	)

	rendered := bldr.Render()

	err := mailer.Send(
		mailer.From,
		[]string{"test@example.com"},
		"Test Email",
		rendered)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Email sent!")
}

// Package mailer provides a simple interface for sending emails.

type Mailer struct {
	Host     string `json:"host,omitempty"`
	Port     int    `json:"port,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	From     string `json:"from,omitempty"`
}

func (m *Mailer) Ready() bool {
	return m.Host != "" && m.Port != 0 && m.Username != "" && m.Password != "" && m.From != ""
}

func (m *Mailer) server() string {
	return m.Host + ":" + strconv.Itoa(m.Port)
}

func (m *Mailer) Send(from string, to []string, subject string, body string) error {
	server := m.server()

	header := map[string]string{
		"From":                      from,
		"Subject":                   mime.QEncoding.Encode("UTF-8", subject),
		"MIME-Version":              "1.0",
		"Content-Type":              "text/html; charset=\"utf-8\"",
		"Content-Transfer-Encoding": "base64",
	}

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(body))

	return smtp.SendMail(
		server,
		smtp.PlainAuth("", m.Username, m.Password, m.Host),
		m.From,
		to,
		[]byte(message),
	)
}
