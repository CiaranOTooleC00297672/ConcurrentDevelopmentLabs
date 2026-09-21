// Ciaran O Toole
// C00297672
// Building a rendezvous using a signal to
// tell other sections that hey im done you can go now
// and that they should wait accordingly.
// Created : 21.09.26

package main

//imports
import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func WorkWithRendezvous(wg *sync.WaitGroup, Num int, signal chan bool) bool {
	var X time.Duration
	X = time.Duration(rand.IntN(5))
	time.Sleep(X * time.Second)
	// decides if it is A or B, A is 0 denoted and B is 1 denoted hence else
	if Num == 0 {
		// no block so A0 runs
		fmt.Println("PartA", Num)
		signal <- true            // send a signal
		<-signal                  // waits for B's signal
		fmt.Println("PartA", Num) // runs once received B0 signal
		signal <- true
	} else {
		<-signal // first is to receive a go ahead
		fmt.Println("PartB", Num)
		signal <- true // sends go ahead
		<-signal       // receive final signal
		fmt.Println("PartB", Num)
	}
	wg.Done()
	return true
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	signal := make(chan bool) // create a channel to say if the process can or cant run basically a switch.
	// both are run using the same wg
	go WorkWithRendezvous(&wg, 0, signal)
	go WorkWithRendezvous(&wg, 1, signal)
	// catches when count hits 0, meaning both of my routines have finished
	wg.Wait()
}

//Notes:
// https://www.youtube.com/watch?v=y2jP45S9BHk - useful video that helped understand fundamentals
// https://www.youtube.com/watch?v=2B-VAxCmhgA - helpful understanding channel communication

//Final thoughts: after countless deadlock issues and burying myself in Youtube videos or online web resources
// my assumption is that this is doing what it is supposed to.
