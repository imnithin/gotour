package main

import (
  "fmt"
)

var (
  sum, a, b int = (1+1), 1, 1
  ToBe float64 = 1.23
)
func main() {
  // fmt.Println(a,b,sum)
  // fmt.Println(ToBe)

    i := 0
    defer fmt.Println(i)
    i++
    return


}
