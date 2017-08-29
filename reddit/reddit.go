package main

import (
	"fmt"
	"log"
  "net/smtp"

	"github.com/Akitsuyoshi/redditnews"
)

func main() {
	to := "YOUR_RECEIVER_ADDRESS"
  subject := "Go articles on Reddit"
  message := redditnewx.Email()

  body := "To: " to + "\r\nSubject: " + subject + "\r\n\r\n" + message
  auth := smtp.PlainAuth("", username, password, "smtp.gmail.com")

  err := smtp.SendMail(
    "smtp.gmail.com:587",
    auth,
    from,
    []string(to),
    [](body),
  )
  if err != nil {
    log.Fatal("SendMail: ", err)
    return
  }
}
