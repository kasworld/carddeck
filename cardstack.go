package carddeck

import (
	"fmt"
	"sync"

	"github.com/kasworld/rand"
)

func NewCardStack() *CardStack {
	return &CardStack{
		cards: make(CardList, 0),
		rnd:   rand.New(),
	}
}

func NewShuffledSingleCardStack(dt *DeckType) *CardStack {
	rtn := NewCardStack()
	rtn.AppendCards(NewCardList(dt))
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
