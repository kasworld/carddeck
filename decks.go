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

//"♠♣♥♦♡♢♤♧☺☻"

var Deck13x4j2 = NewDeckType(
	[]string{"♠", "♡", "♢", "♣"},
	[]string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"},
	[]string{"☻", "☺"},
)
var Deck13x4j1 = NewDeckType(
	[]string{"♠", "♡", "♢", "♣"},
	[]string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"},
	[]string{"☻"},
)

var Deck13x4 = NewDeckType(
	[]string{"♠", "♡", "♢", "♣"},
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
