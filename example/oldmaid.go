package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"

	"github.com/kasworld/carddeck"
	"github.com/kasworld/log"
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
	name string
	carddeck.CardList
	rnd *rand.Rand
}

func NewPlayer(name string) *Player {
	return &Player{
		name: name,
		rnd:  rand.New(),
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
	return fmt.Sprintf("%v:%v", p.name, p.CardList)
}

type OldMaidGame struct {
	players   []*Player
	discarded carddeck.CardList
	played    []*Player
	rnd       *rand.Rand
}

func NewGame(n int) *OldMaidGame {
	players := make([]*Player, n)
	for i, _ := range players {
		players[i] = NewPlayer(fmt.Sprintf("player%d", i))
	}
	return &OldMaidGame{
		players:   players,
		discarded: make(carddeck.CardList, 0),
		played:    make([]*Player, 0),
		rnd:       rand.New(),
	}
}
func (g *OldMaidGame) Init() {
	// prepare playing deck
	cl := carddeck.NewCardList(carddeck.Deck13x4)
	cl.DrawByPos(g.rnd.Intn(len(cl)))
	cs := carddeck.NewCardStack()
	cs.AppendCards(cl)
	cs.Shuffle()

	// distribute card to player
	for i := 0; ; i++ {
		cd := cs.DrawCard()
		if cd == nil {
			break
		}
		g.players[i%len(g.players)].Append(cd)
	}
	for _, v := range g.players {
		fmt.Printf("%v", v)
		v.removeSameNum(&g.discarded)
		fmt.Printf(" %v\n", v)
	}
}

func (g *OldMaidGame) Print() {
	for _, v := range g.players {
		v.SortNum()
		fmt.Printf("%v\n", v)
	}
}

func (g *OldMaidGame) RemovePlayer(i int) {
	if i >= len(g.players) {
		log.Warn("invalid player number %v %v", i, len(g.players))
		return
	}
	g.played = append(g.played, g.players[i])
	g.players = append(g.players[:i], g.players[i+1:]...)
}

func (g *OldMaidGame) Step() {
	for i := 0; i < len(g.players); i++ {
		srcplayer := g.players[i]
		dstpos := (i + 1) % len(g.players)
		dstplayer := g.players[dstpos]
		srcplayer.Draw1Card(dstplayer)
		fmt.Printf("%v %v\n", srcplayer, dstplayer)
		if len(dstplayer.CardList) == 0 {
			// log.Info("player end %v", dstplayer)
			g.RemovePlayer(dstpos)
			i++
		}
		srcplayer.removeSameNum(&g.discarded)
		if len(srcplayer.CardList) == 0 {
			// log.Info("player end %v", srcplayer)
			g.RemovePlayer(i)
		}
		// fmt.Printf("%v:%v ", i, srcplayer)
	}
	// fmt.Println()
}

func doMain(args []string) {
	g := NewGame(6)
	g.Init()
	g.Print()
	for len(g.players) > 1 {
		g.Step()
	}
	g.RemovePlayer(0)
	fmt.Printf("%v\n", g.discarded)
	for _, v := range g.played {
		fmt.Printf("%v\n", v)
	}
}
