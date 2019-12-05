package printer

import (
	printSpec "gopra/delegate/printspec"
	"log"
)

// Printer 列印者
type Printer struct {
	R printSpec.IPrintCore
}

// PrintSomeThing 印些東西
func (p *Printer) PrintSomeThing() {
	p.R.ReceiveAndPrintHandler(func(s string) {
		log.Println("get something to print : ", s)
	})
}
