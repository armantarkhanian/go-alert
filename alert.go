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

func (a *AlertInstance) Send(message string) error {
	for _, alerter := range a.alertManagers {
		if err := alerter.Alert(message); err != nil {
			return err
		}
	}

	return nil
}
