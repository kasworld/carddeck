package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"

	"github.com/kasworld/carddeck"
	// "github.com/kasworld/rand"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
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

func doMain(args []string) {
	c := carddeck.NewCards(carddeck.Deck13x4j2)
	// c.Shuffle(rand.New())
	// c.Sort()
	cs := carddeck.NewCardStack()
	cs.AppendCards(c)
	cs.Shuffle()

	hands := make([]carddeck.CardList, 4)

	for i := 0; ; i++ {
		cd := cs.DrawCard()
		if cd == nil {
			break
		}
		hands[i%4].Append(cd)
	}
	for i, v := range hands {
		v.SortNum()
		fmt.Printf("%v:%v\n", i, v)
	}
}
