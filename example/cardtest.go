package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"

	"github.com/kasworld/carddeck"
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
	cs := carddeck.NewCardStack()
	cs.AppendCards(c)
	cs.Shuffle()

	for cd := cs.DrawCard(); cd != carddeck.CardEmpty; cd = cs.DrawCard() {
		fmt.Printf("%v\n", cd)
	}
}
