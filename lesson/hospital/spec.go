package hospital

// IHostpital spec
type IHostpital interface {
	CheckInfected(func(map[string]bool))
	SentPatient(func([]string))
}
