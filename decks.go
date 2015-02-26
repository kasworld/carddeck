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
