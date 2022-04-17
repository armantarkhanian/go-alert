package alert

func (a *AlertInstance) Custom(alerter AlertManager) *AlertInstance {
	a.alertManagers = append(a.alertManagers, alerter)
	return a
}
