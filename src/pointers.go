package main

import "fmt"

func main() {
  i, j := 42, 2701

  p := &i
  *p = 21

  fmt.Println(i)
  fmt.Println(j)
}
