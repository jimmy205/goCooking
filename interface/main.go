package main

// ICore i core
type ICore interface {
	Deal()
}

// Core core
type Core struct {
	cards []string
}

// War war
type War struct {
	c *Core
}

// IWar i war
type IWar interface {
}

// NewShuffler new shuffler
func NewShuffler() *Core {
	return &Core{}
}

// NewCardSet new card set
func (c *Core) NewCardSet() {
	cards := []string{"a", "b", "c"}
	c.cards = cards
}

// NewWar new war
func NewWar() *War {
	c := NewShuffler()

	return &War{
		c: c,
	}

}

// Deal deal
func (w *War) Deal() {

}

// JudgeWin judge win
func (w *War) JudgeWin() {

}

// NewGame new game
func NewGame(t string) ICore {

	if t == "war" {
		return NewWar()
	}

	return nil
}

func main() {

}
