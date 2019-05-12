package Send_Email

import (
	"errors"
	"fmt"
	"net/smtp"
	"strings"
)

type Message struct {
	From    string
	To      []string
	Cc      []string
	Subject string
	Content string
}

func (m Message) Bytes() (r []byte) {
	to := strings.Join(m.To, ",")
	cc := strings.Join(m.Cc, ",")

	r = append(r, []byte("From: "+m.From+"\n")...)
	r = append(r, []byte("To: "+to+"\n")...)
	r = append(r, []byte("Cc: "+cc+"\n")...)
	r = append(r, []byte("Subject: "+m.Subject+"\n\n")...)
	r = append(r, []byte(m.Content)...)

	return
}

func (m Message) Send(host string, port int, user, pass string) (err error) {
	err = check(host, user, pass)
	if err != nil {
		return
	}

	err = smtp.SendMail(fmt.Sprintf("%v:%v", host, port),
		smtp.PlainAuth("", user, pass, host),
		m.From,
		m.To,
		m.Bytes(),
	)

	return
}

func Trim(input string) string {
	return strings.Trim(input, " \t\n\r")
}

func check(host, user, pass string) error {
	if host == "" {
		return errors.New("bad host")
	}
	if user == "" {
		return errors.New("bad username")
	}
	if pass == "" {
		return errors.New("bad password")
	}

	return nil
}
