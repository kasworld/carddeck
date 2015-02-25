package carddeck

import (
	"fmt"
	"sort"
	"sync"

	"github.com/kasworld/rand"
)

var Deck13x4j2 = NewDeckType(
	[]string{"Spades", "Hearts", "Diamonds", "Clubs"},
	[]string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"},
	[]string{"Major-Joker", "minor-joker"},
)
var Deck13x4j1 = NewDeckType(
	[]string{"Spades", "Hearts", "Diamonds", "Clubs"},
	[]string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"},
	[]string{"Joker"},
)

var Deck13x4 = NewDeckType(
	[]string{"Spades", "Hearts", "Diamonds", "Clubs"},
	[]string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"},
	nil,
)
var DeckRWTarot = NewDeckType(
	[]string{"Sword", "Cup", "Wand", "Pentacle"},
	[]string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Page", "Knight", "Queen", "King"},
	[]string{"0–The Fool",
		"I–The Magician",
		"II–The High Priestess",
		"III–The Empress",
		"IV–The Emperor",
		"V–The Hierophant",
		"VI–The Lovers",
		"VII–The Chariot",
		"VIII–Strength",
		"IX–The Hermit",
		"X–Wheel of Fortune",
		"XI–Justice",
		"XII–The Hanged Man",
		"XIII–Death",
		"XIV–Temperance",
		"XV–The Devil",
		"XVI–The Tower",
		"XVII–The Star",
		"XVIII–The Moon",
		"XIX–The Sun",
		"XX–Judgement",
		"XXI–The World"},
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

func NewCards(dt *DeckType) CardList {
	rtn := make(CardList, dt.suitCount*dt.numCount+dt.jokerCount)
	for i := 0; i < dt.suitCount*dt.numCount+dt.jokerCount; i++ {
		rtn[i] = &Card{i, dt}
	}
	return rtn
}

func NewCardStack() *CardStack {
	return &CardStack{
		cards: make(CardList, 0),
		rnd:   rand.New(),
	}
}

func NewShuffledSingleCardStack(dt *DeckType) *CardStack {
	rtn := NewCardStack()
	rtn.AppendCards(NewCards(dt))
	rtn.Shuffle()
	return rtn
}

type CardStack struct {
	cards CardList
	rnd   *rand.Rand
	pos   int
	mutex sync.Mutex
}

func (cs CardStack) String() string {
	return fmt.Sprintf("%v", cs.cards)
}

func (cs *CardStack) AppendCards(ncs CardList) {
	cs.mutex.Lock()
	defer cs.mutex.Unlock()
	cs.cards = append(cs.cards, ncs...)
}

func (cs *CardStack) rewind() {
	cs.pos = 0
}

func (cs *CardStack) Rewind() {
	cs.mutex.Lock()
	defer cs.mutex.Unlock()
	cs.rewind()
}
func (cs *CardStack) Shuffle() {
	cs.mutex.Lock()
	defer cs.mutex.Unlock()
	cs.cards.Shuffle(cs.rnd)
	cs.rewind()
}

func (cs *CardStack) DrawCard() *Card {
	cs.mutex.Lock()
	defer cs.mutex.Unlock()
	if cs.Empty() {
		return nil
	}
	rtn := cs.cards[cs.pos]
	cs.pos++
	return rtn
}

func (cs *CardStack) Empty() bool {
	return cs.pos >= len(cs.cards)
}
