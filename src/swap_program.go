package main

import(
  "fmt"
)

func swap(x string, y string) (string, string) {
  return y, x
}

func main() {
  a, b := swap("k", "nithin")
  fmt.Printf("a and b: %s %s\n", a, b)
  fmt.Println(swap(a,b))
}
