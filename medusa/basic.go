package medusa

// Medusa 設定檔
type Medusa struct {
	baseURL   string
	retry     int
	timeout   int
	useCache  bool
	Envirment Envirment
}

// Envirment 環境
type Envirment struct {
	Local       string
	Development string
	QaTest      string
	Production  string
}

// New 新建一個medusa
func New() *Medusa {
	return &Medusa{
		retry:    3,
		timeout:  3,
		useCache: true,
		Envirment: Envirment{
			Local:       "local",
			Development: "development",
			QaTest:      "qaTest",
			Production:  "production",
		},
	}
}

// SetDomain 設定Domain
func (m *Medusa) SetDomain(env string) {

	var (
		http   string
		domain string
	)

	http = "http://"

	switch env {
	case m.Envirment.Local:
		domain = "127.0.0.1"
	case m.Envirment.Development:
		domain = "mdu.xbb-slot-dev.com"
	case m.Envirment.QaTest:
		http = "https://"
	case m.Envirment.Production:
		http = "https://"
	}

	m.baseURL = http + domain
}

// SetTimeout 設定過期時間
func (m *Medusa) SetTimeout(n int) {
	m.timeout = n
}

// SetRetry 設定重試
func (m *Medusa) SetRetry(n int) {
	m.retry = n
}

// SetUserCache 設定是否使用快取
func (m *Medusa) SetUserCache(b bool) {
	m.useCache = b
}
