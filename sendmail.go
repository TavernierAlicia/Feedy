package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"

	gomail "gopkg.in/gomail.v2"
)

// RecvMail
func recvMail(mail string, name string, message string) (err error) {

	to := "servicefeedy@gmail.com"

	//configure sending mailbox
	from := "servicefeedy@gmail.com"
	pass := ""

	subject := "Feedy, question from " + name + " (" + mail + ")"
	//set message
	message = " \n" + message

	msg := "From: " + mail + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		message

	err = smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		fmt.Println("smtp error: %s", err)
		fmt.Println("failed")
	} else {
		fmt.Println("passed")
	}
	return err
}

//sendmail
func sendMail(mail string, name string, message string) (err error) {

	t := template.New("mailing.html")

	t, err = t.ParseFiles("mailing.html")
	if err != nil {
		fmt.Println(err)
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, ""); err != nil {
		fmt.Println(err)
	}

	result := tpl.String()
	m := gomail.NewMessage()
	m.SetHeader("From", "servicefeedy@gmail.com")
	m.SetHeader("To", mail)
	m.SetAddressHeader("Cc", mail, mail)
	m.SetHeader("Subject", "Découvrez nos offres !")
	m.SetBody("text/html", result)
	m.Attach("Feedy_PDF.pdf")

	d := gomail.NewDialer("smtp.gmail.com", 587, "servicefeedy@gmail.com", "")

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return err
}
