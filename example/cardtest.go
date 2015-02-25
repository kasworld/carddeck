package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"

	"github.com/kasworld/carddeck"
	"github.com/kasworld/rand"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var profilefilename = flag.String("pfilename", "", "profile filename")
	flag.Parse()
	args := flag.Args()

	if *profilefilename != "" {
		f, err := os.Create(*profilefilename)
		if err != nil {
			log.Fatalf("profile %v", err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	doMain(args)
}

type Player struct {
	carddeck.CardList
	rnd *rand.Rand
}

func NewPlayer() *Player {
	return &Player{
		rnd: rand.New(),
	}
}

func (p *Player) removeSameNum(ds *carddeck.CardList) {
	p.SortNum()
	for i := len(p.CardList) - 1; i > 0; i-- {
		if p.CardList[i].Num() == p.CardList[i-1].Num() {
			c := p.DrawByPos(i)
			ds.Append(c)
			c = p.DrawByPos(i - 1)
			ds.Append(c)
			i--
		}
	}
	// p.Shuffle(p.rnd)
}

func (p *Player) Draw1Card(p2 *Player) {
	pos := p.rnd.Intn(len(p2.CardList))
	p.Append(p2.DrawByPos(pos))
}

func (p Player) String() string {
	return fmt.Sprintf("%v", p.CardList)
}

func doMain(args []string) {

	// prepare playing deck
	cs := carddeck.NewShuffledSingleCardStack(carddeck.Deck13x4j1)
	discarded := make(carddeck.CardList, 0)

	// make player
	players := make([]*Player, 4)
	for i, _ := range players {
		players[i] = NewPlayer()
	}

	// distribute card to player
	for i := 0; ; i++ {
		cd := cs.DrawCard()
		if cd == nil {
			break
		}
		players[i%len(players)].Append(cd)
	}

	// player arrange hands
	for i, v := range players {
		v.SortNum()
		fmt.Printf("%v:%v\n", i, v)
	}
	for i, v := range players {
		v.removeSameNum(&discarded)
		fmt.Printf("%v:%v\n", i, v)
	}

}
