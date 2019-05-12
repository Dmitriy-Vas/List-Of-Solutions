package Send_Email

import (
	"testing"
)

// Before running tests
// Specify your credentials and mail
func TestSend(t *testing.T) {
	data := struct {
		host string
		port int
		user string
		pass string
	}{
		host: "",
		port: 587,
		user: "",
		pass: "",
	}

	from := ""
	from = Trim(from)

	to := []string{"", "", ""}
	for i, t := range to {
		to[i] = Trim(t)
	}

	cc := []string{"", "", ""}
	for i, c := range cc {
		cc[i] = Trim(c)
	}

	subject := ""
	subject = Trim(subject)

	content := ""
	content = Trim(content)

	m := &Message{
		From:    from,
		To:      to,
		Cc:      cc,
		Subject: subject,
		Content: content,
	}

	if err := m.Send(data.host, data.port, data.user, data.pass); err != nil {
		t.Fatal(err)
	}
}
