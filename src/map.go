package main
import "fmt"

type Vertex struct {
  x float64
  y float64
}

var m map[string]Vertex

func main(){
  m = make(map[string]Vertex)
  m["Bill Clinton"] = Vertex{
    40.0, -10.93,
  }
  fmt.Println(m["Bill Clinton"])

}
