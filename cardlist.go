package carddeck

import (
	"sort"

	"github.com/kasworld/rand"
)

type By func(p1, p2 *Card) bool

func (by By) Sort(cards []*Card) {
	ps := &cardSorter{
		cards: cards,
		by:    by, // The Sort method's receiver is the function (closure) that defines the sort order.
	}
	sort.Sort(ps)
}

type cardSorter struct {
	cards []*Card
	by    func(p1, p2 *Card) bool // Closure used in the Less method.
}

func (s *cardSorter) Len() int {
	return len(s.cards)
}
func (s *cardSorter) Swap(i, j int) {
	s.cards[i], s.cards[j] = s.cards[j], s.cards[i]
}
func (s *cardSorter) Less(i, j int) bool {
	return s.by(s.cards[i], s.cards[j])
}

func (s CardList) SortSuit() {
	suits := func(p1, p2 *Card) bool {
		return p1.v < p2.v
	}
	By(suits).Sort(s)
}
func (s CardList) SortNum() {
	suits := func(p1, p2 *Card) bool {
		if p1.IsJoker() && !p2.IsJoker() {
			return false
		}
		if !p1.IsJoker() && p2.IsJoker() {
			return true
		}
		if p1.Num() == p2.Num() {
			return p1.Suit() < p2.Suit()
		}
		return p1.Num() < p2.Num()
	}
	By(suits).Sort(s)
}

type CardList []*Card

func (s CardList) FindIndex(v int) int {
	return sort.Search(len(s), func(i int) bool { return s[i].v >= v })
}
func (s CardList) Shuffle(rnd *rand.Rand) {
	n := len(s)
	for i := 0; i < n; i++ {
		j := rnd.Intn(i + 1)
		s[i], s[j] = s[j], s[i]
	}
}
func (s *CardList) Append(c *Card) {
	*s = append(*s, c)
}
func (s CardList) Add(cs CardList) CardList {
	return append(s, cs...)
}
func (s *CardList) DrawByPos(i int) *Card {
	if i >= len(*s) {
		return nil
	}
	rtn := (*s)[i]
	*s = append((*s)[:i], (*s)[i+1:]...)

	return rtn
}
func NewCardList(dt *DeckType) CardList {
	rtn := make(CardList, dt.suitCount*dt.numCount+dt.jokerCount)
	for i := 0; i < dt.suitCount*dt.numCount+dt.jokerCount; i++ {
		rtn[i] = &Card{i, dt}
	}
	return rtn
}
