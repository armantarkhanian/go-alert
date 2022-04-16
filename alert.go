package alert

type AlertManager interface {
	Alert(message string) error
}

type AlertInstance struct {
	alerters []AlertManager
}

func (a *AlertInstance) Send(message string) error {
	for _, alerter := range a.alerters {
		if err := alerter.Alert(message); err != nil {
			return err
		}
	}

	return nil
}

func New(alerters ...AlertManager) *AlertInstance {
	return &AlertInstance{
		alerters: alerters,
	}
}
