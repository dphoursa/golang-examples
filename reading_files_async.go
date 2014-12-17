package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "runtime"
  "sync"
)

func readLines(path string) error {
  file, err := os.Open(path)
  if err != nil {
    return err
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    fmt.Println(scanner.Text())
  }
  return scanner.Err()
}

func readFile (wg *sync.WaitGroup, filename string) {
  defer wg.Done()
  
  err := readLines(filename)
  if err != nil {
    log.Println("readLines: %s", err)
  }
}

func main() {
  runtime.GOMAXPROCS(2)
  args := os.Args[1:]
  
  var wg sync.WaitGroup
  
  for i := range args {
    wg.Add(1)
    go readFile(&wg, args[i])
  }
  
  wg.Wait()
}