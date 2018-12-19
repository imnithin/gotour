package main

import (
  "fmt"
  "math"
  "strconv"
)
func sqrt(x float64) string {
  if x < 0 {
    return sqrt(-x) + "i"
  }
  return fmt.Sprint(math.Sqrt(x))
  return strconv.FormatFloat(math.Sqrt(x), 'f', 1, 64)
}

func main() {
  fmt.Println(sqrt(2), sqrt(-4))
}
