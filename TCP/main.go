package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
  ports := make(chan int, 100)
  var wg sync.WaitGroup
  for i := 1: i < 1024; i++ {
    go worker(ports, &wg)
    
  }
  for i := 1; i < 1024; i++ {
    wg.Add(1)
    ports <- i
    
  }
  wg.Wait()
  close(ports)
}

func worker(ports chan int, wg *syn.WaitGroup){
  for p := range ports
  fmt.Println(p)
  wg.Done()
}
