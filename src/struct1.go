package main

import (
  "fmt"
)
func main() {

  arr1 := []int{2, 3, 4}
  arr2 := []bool{true, false, true, true}
  fmt.Println(arr1)
  fmt.Println(arr2)

  struct1 := []struct{
    i int
    z bool
  } {
      {0, false},
      {1, true},
  }
  fmt.Println(struct1)

}
