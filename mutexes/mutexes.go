package main

import (
	"fmt"
	"sync"
	"time"
)

type safeCounter struct {
	counts map[string]int
	mux    *sync.Mutex
	mux2   *sync.RWMutex
}

func (sc safeCounter) inc(key string) {
	sc.mux.Lock()
	defer sc.mux.Unlock()
	sc.slowIncrement(key)
}
func (sc safeCounter) val(key string) int {
	sc.mux.Lock()
	defer sc.mux.Unlock()
	return sc.counts[key]
}
func (sc safeCounter) slowIncrement(key string) {
	tempCounter := sc.counts[key]
	time.Sleep(time.Microsecond)
	tempCounter++
	sc.counts[key] = tempCounter
}

// RW Mutex(it can give you read and write locak and unlock also along with lock and unlock)
func (sc safeCounter) rval(key string) int {
	sc.mux2.RLock()
	defer sc.mux2.RUnlock()
	return sc.counts[key]
}

func main() {
	sc := safeCounter{
		counts: make(map[string]int),
		mux:    &sync.Mutex{},
	}
	fmt.Println(sc.counts["sharmakeshav54126@gmail.com"])
	sc.inc("keshav.sharma@novostack.com")
	val := sc.val("keshav.sharma@novostack.com")
	fmt.Println(val)
}
