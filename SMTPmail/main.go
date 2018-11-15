/* ========================================================================
 *  File: main.go
 *  Date: 2018-11-15
 *  Creator: DimiG
 * ========================================================================
 */

package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/smtp"
	"os"
)

func send(body string, host string, from string, to string, pass string) {
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Message from user...\n\n" +
		body

	err := smtp.SendMail(host+":587",
		smtp.PlainAuth("", from, pass, host),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("SMTP ERROR: %s", err)
		return
	}

	log.Println("Mail Sent... ;)")
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file. Using the SYSTEM one.")
	}

	var (
		MAIL_ADDR_FROM = os.Getenv("MAIL_ADDR_FROM")
		MAIL_ADDR_TO   = os.Getenv("MAIL_ADDR_TO")
		MAIL_PWD       = os.Getenv("MAIL_PWD")
		MAIL_HOST      = os.Getenv("MAIL_HOST")
	)

	send("HELLO GUYs !!! :)", MAIL_HOST, MAIL_ADDR_FROM, MAIL_ADDR_TO, MAIL_PWD)
}
