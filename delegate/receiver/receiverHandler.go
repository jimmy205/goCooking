package receiver

// ReceiveAndPrintHandler 通知可以列印
func (r *Receiver) ReceiveAndPrintHandler(fn func(string)) {
	r.rHandlePrint = fn
}
