package carddeck

type Card struct {
	v    int
	deck *DeckType
}

func (c Card) IsJoker() bool {
	return c.deck.IsJoker(c.v)
}
func (c Card) Suit() int {
	return c.deck.Suit(c.v)
}
func (c Card) Num() int {
	return c.deck.Num(c.v)
}
func (c Card) String() string {
	return c.deck.ToStr(c.v)
}
