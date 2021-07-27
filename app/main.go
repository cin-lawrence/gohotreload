package main

import (
  "log"
  "time"
)

func main() {
  for {
    log.Println("Hello, world")
    time.Sleep(2 * time.Second)
  }
}
