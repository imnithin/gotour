package main
import "fmt"

// type Vertex struct {
//   x int
//   y int
// }
// func main() {
//   v := Vertex{100,22}
//   v.x = 121
//   fmt.Println(v.x)
// }

func main() {
  k := []struct {
     a int
     z bool
  }{
    {1,true},
    {2,true},
  }

  fmt.Println(k)

}
