// Copyright 2015 SeukWon Kang (kasworld@gmail.com)
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
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
