package carddeck

import (
	"fmt"
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

var CardEmpty = Card{-1, nil}

func (c Card) Suit() int {
	return c.deck.Suit(c.v)
}
func (c Card) Num() int {
	return c.deck.Num(c.v)
}
func (c Card) String() string {
	return c.deck.ToStr(c.v)
}

func NewCards(dt *DeckType) []Card {
	rtn := make([]Card, dt.suitCount*dt.numCount+dt.jokerCount)
	for i := 0; i < dt.suitCount*dt.numCount+dt.jokerCount; i++ {
		rtn[i] = Card{i, dt}
	}
	return rtn
}

func NewCardStack() *CardStack {
	return &CardStack{
		cards: make([]Card, 0),
		rnd:   rand.New(),
	}
}

type CardStack struct {
	cards []Card
	rnd   *rand.Rand
	pos   int
	mutex sync.Mutex
}

func (cs CardStack) String() string {
	return fmt.Sprintf("%v", cs.cards)
}

func (cs *CardStack) AppendCards(ncs []Card) {
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
	n := len(cs.cards)
	for i := 0; i < n; i++ {
		j := cs.rnd.Intn(i + 1)
		cs.cards[i], cs.cards[j] = cs.cards[j], cs.cards[i]
	}
	cs.rewind()
}

func (cs *CardStack) DrawCard() Card {
	cs.mutex.Lock()
	defer cs.mutex.Unlock()
	if cs.Empty() {
		return CardEmpty
	}
	rtn := cs.cards[cs.pos]
	cs.pos++
	return rtn
}

func (cs *CardStack) Empty() bool {
	return cs.pos >= len(cs.cards)
}
