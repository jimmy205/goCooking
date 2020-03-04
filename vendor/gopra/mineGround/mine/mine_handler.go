package mine

// GoldHandler 礦場提供金礦產生的Func
func (g *Ground) GoldHandler(fn func(string)) {
	g.gGoldHandler = fn
}

// RockHandler 礦場提供石頭產生的Func
func (g *Ground) RockHandler(fn func(string)) {
	g.gRockHandler = fn
}
