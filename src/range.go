package main

import "fmt"

func main(){

  var arr = []int{2,3,5,6}

  for i,v := range arr {
    fmt.Printf("%d, %d\n",i,v)
  }
}
