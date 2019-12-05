package printspec

// IPrintCore 影印的核心
type IPrintCore interface {
	ReceiveAndPrintHandler(func(string))
}
