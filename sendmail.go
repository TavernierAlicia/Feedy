package main

import (
	"fmt"
	"net/smtp"
)

//sendmail
func SendMail(mail string, name string, message string) (err error) {

	to := "tavernieralicia00@gmail.com"

	//configure sending mailbox
	from := "tavernieralicia00@gmail.com"
	pass := ""

	subject := "Feedy, question from "+name+" - "+mail
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