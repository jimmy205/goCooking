package hotel

// ReceivePatient 接收病患
func (m *Manager) ReceivePatient() {
	m.hospital.SentPatient(func(nameList []string) {
		for _, name := range nameList {
			m.JoinRoom(NewTourist(name, false))
		}
	})
}

// ReceiveCheckResult 接收病患檢查結果
func (m *Manager) ReceiveCheckResult() {
	m.hospital.CheckInfected(func(result map[string]bool) {
		for r := range m.roomsMap {
			for t := range r.touristMap {
				if result[t.name] {
					continue
				}
				r.CheckOut(t)
			}
		}
	})
}
