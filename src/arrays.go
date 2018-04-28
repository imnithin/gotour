package main

import "fmt"

func main() {
  var a [2]string
  a[0] = "Hello"
  a[1] = "World"
  fmt.Println(a)
  fmt.Println(a[0], a[1])
  // fmt.Println(a[2])
  // var primes [6]int
  primes := [6]int{2,3,5,7,11,13}
  fmt.Println(primes)
  var s []int := primes[1:2]
  fmt.Println(s)
}
