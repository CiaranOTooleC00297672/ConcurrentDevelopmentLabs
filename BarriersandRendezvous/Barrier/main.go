// Ciaran O Toole
// C00297672
// Building a barrier using semaphores, mutex locks.
// Created : 21.09.26

package main

import (
	"fmt"
	"sync"
	"time"
)

// define our semaphore structure
type semaphore struct {
	theCounter chan struct{}
}

// Place a barrier in this function --use Mutex's and Semaphores
func doStuff(goNum int, count *int, wg *sync.WaitGroup, sem chan struct{}, tot *int, mutex *sync.Mutex) bool {
	time.Sleep(time.Second)
	fmt.Println("Part A", goNum)
	mutex.Lock()        //we lock here as if we let them all access our counter at the same time it can break (dangerous)
	*count++            //increment process amount, we use mem address and pointers so counter is consistent
	if *count == *tot { //if we hit 10, that's all of A so last one lets them all run then we allow B to work
		for i := 0; i < *tot; i++ {
			sem <- struct{}{} //cycle for 10 and output them
		}
	}
	mutex.Unlock()              //unlock
	<-sem                       //if channels empty it blocks if not itll allow run
	fmt.Println("PartB", goNum) //B runs
	wg.Done()
	return true
}

func main() {
	totalRoutines := 10                              //total we want of both A and B
	total := 10                                      // a max num for tracking all A
	count := 0                                       // counter to compare
	semaphore2 := make(chan struct{}, totalRoutines) //semaphore to handle the process count
	var wg sync.WaitGroup
	wg.Add(totalRoutines)
	//we will need some of these
	var mutex sync.Mutex //our mutex defined

	for i := range totalRoutines { //create the go Routines here
		go doStuff(i, &count, &wg, semaphore2, &total, &mutex) //run function
	}
	wg.Wait() //wait for everyone to finish before exiting
}

//People who helped me: Bartosz Liberda
//People I helped: Sam Geraghty
// Bartosz helped give the idea of the semaphore to me so i went with that in comparison to how he wrote it
