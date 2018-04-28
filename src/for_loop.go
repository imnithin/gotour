package main

import (
  "fmt"
)

func main() {
  // sum := 0
  var sum int = 0
  for i := 0; i < 100; i++ {
    sum += i
  }
  fmt.Println(sum)
}
