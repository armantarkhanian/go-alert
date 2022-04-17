package alert

type AlertManager interface {
	Alert(message string) error
}

type AlertInstance struct {
	alertManagers []AlertManager
}

func New() *AlertInstance {
	return &AlertInstance{}
}

// Telegram adds telegram support
func (a *AlertInstance) Telegram(botToken string, chatIDs []int64) *AlertInstance {
	if len(chatIDs) == 0 {
		return a
	}

	a.alertManagers = append(a.alertManagers, &TelegramAlerter{
		chatIDs:  chatIDs,
		botToken: botToken,
	})

	return a
}

// Email adds email support
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

// Custom adds support for custom implementation of AlertManager interface
func (a *AlertInstance) Custom(alerter AlertManager) *AlertInstance {
	a.alertManagers = append(a.alertManagers, alerter)
	return a
}

// Send sends the message to all AlertManagers
func (a *AlertInstance) Send(message string) error {
	for _, am := range a.alertManagers {
		if err := am.Alert(message); err != nil {
			return err
		}
	}

	return nil
}
