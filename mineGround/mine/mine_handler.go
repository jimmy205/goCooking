package mine

// GoldHandler 通知金礦產生
func (g *Ground) GoldHandler(fn func(string)) {
	g.gGoldHandler = fn
}

// RockHandler 通知石頭產生
func (g *Ground) RockHandler(fn func(string)) {
	g.gRockHandler = fn
}
