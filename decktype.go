package carddeck

import (
	"fmt"
)

type DeckType struct {
	suitCount  int
	numCount   int
	jokerCount int
	suitRunes  []string
	numRunes   []string
	jokerRunes []string
}

func NewDeckType(s, n, j []string) *DeckType {
	return &DeckType{
		len(s), len(n), len(j),
		s, n, j,
	}
}
func (dt *DeckType) IsJoker(v int) bool {
	return dt.Suit(v) >= dt.suitCount
}
func (dt *DeckType) Suit(v int) int {
	s := v / dt.numCount
	return s
}
func (dt *DeckType) Num(v int) int {
	if dt.IsJoker(v) {
		return v - dt.suitCount*dt.numCount
	} else {
		n := v % dt.numCount
		return n
	}
}
func (dt *DeckType) ToStr(v int) string {
	s := dt.Suit(v)
	n := dt.Num(v)
	if s < dt.suitCount {
		return fmt.Sprintf("%v-%v", dt.suitRunes[s], dt.numRunes[n])
	} else {
		return fmt.Sprintf("%v", dt.jokerRunes[n])
	}
}
