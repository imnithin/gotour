package main
import "fmt"

type Vertex struct {
  Lat float64
  Long float64
}

var m = map[string]Vertex {
  "bill" : Vertex{1,2},
  "gates" : Vertex{1,2},
}

func main() {
  fmt.Println(m)
}
