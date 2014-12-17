package main

import (
  "fmt"
  "os"
  "runtime"
  "sync"
)

func printArg (wg *sync.WaitGroup, val string) {
  defer wg.Done()
  
  fmt.Println(val)
}

func main() {
  runtime.GOMAXPROCS(2)
  args := os.Args[1:]
  
  var wg sync.WaitGroup
  
  for i := range args {
    wg.Add(1)
    go printArg(&wg, args[i])
  }
  
  wg.Wait()
}