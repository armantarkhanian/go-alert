package alert

import (
	"net"
	"net/smtp"
)

type EmailAlerter struct {
	host     string
	from     string
	password string
	to       []string
}

var _ AlertManager = &EmailAlerter{}

func (t *EmailAlerter) Alert(message string) error {
	host, _, err := net.SplitHostPort(t.host)
	if err != nil {
		return err
	}

	msg := "Subject: Alert\r\n\r\n" + message + "\r\n"

	return smtp.SendMail(
		t.host,
		smtp.PlainAuth("", t.from, t.password, host),
		t.from,
		t.to,
		[]byte(msg),
	)
}

func (a *AlertInstance) Email(host, from, password string, to []string) *AlertInstance {
	if len(to) == 0 {
		return a
	}

	a.alertManagers = append(a.alertManagers, &EmailAlerter{
		host:     host,
		from:     from,
		password: password,
		to:       to,
	})
	return a
}
