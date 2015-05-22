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
