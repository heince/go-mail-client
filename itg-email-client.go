package main

import (
	"flag"
	"github.com/scorredoira/email"
	"io/ioutil"
	"log"
	"net/mail"
	"net/smtp"
)

func main() {

	attachmentPtr := flag.String("attach", "empty", "-attach [attachment file path]")
	fromPtr := flag.String("from", "empty", "-from=[email address]")
	toPtr := flag.String("to", "empty", "-to=[email address]")
	subjectPtr := flag.String("subject", "empty", "-subject=[subject name]")
	bodyPtr := flag.String("body", "empty", "-body=[body file path]")
	contentPtr := flag.String("content", "text", "-content=[text/html] (default to text)")
	usernamePtr := flag.String("user", "", "-user=[username]")
	passwordPtr := flag.String("password", "", "-password=[password]")
	hostPtr := flag.String("host", "localhost", "-host=[email server] (default to localhost)")
	portPtr := flag.Uint("port", 25, "-port=[port number] (default to 25)")

	flag.Parse()

	if *fromPtr == "empty" {
		log.Fatal("-from [email address] must be specified")
	}

	if *toPtr == "empty" {
		log.Fatal("-to [email address] must be specified")
	}

	if *bodyPtr == "empty" {
		log.Fatal("-body must be specified")
	}

	body, err := ioutil.ReadFile(*bodyPtr)
	check(err)
	bodyString := string(body)

	var m *email.Message

	if *contentPtr == "text" {
		m = email.NewMessage(*subjectPtr, bodyString)
	} else {
		m = email.NewHTMLMessage(*subjectPtr, bodyString)
	}

	m.From = mail.Address{Address: *fromPtr}
	m.To = []string{*toPtr}

	if *attachmentPtr != "empty" {
		if err := m.Attach(*attachmentPtr); err != nil {
			log.Fatal(err)
		}
	}

	auth := smtp.PlainAuth("", *usernamePtr, *passwordPtr, *hostPtr)

	hostPort := *toPtr + ":" + string(*portPtr)
	if err := email.Send(hostPort, auth, m); err != nil {
		log.Fatal(err)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
