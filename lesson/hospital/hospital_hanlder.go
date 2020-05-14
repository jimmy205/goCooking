package hospital

// SentPatient 傳送新發現的病患
func (h *Hospital) SentPatient(fn func([]string)) {
	h.sendPatientHandler = fn
}

// CheckInfected 確認這個病患有沒有生病
func (h *Hospital) CheckInfected(fn func(map[string]bool)) {
	h.sendPatientTestResult = fn
}
