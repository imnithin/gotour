package main
import "fmt"


type Vertex struct {
  X, Y float64
}

func Scale(v *Vertex, f float64) {
  v.X = v.X * f
  v.Y = v.Y * f
}

func (v *Vertex) ScaleFunc(f float64){
  v.X = v.X * f
  v.Y = v.Y * f
}

func main() {
  v := Vertex{3,4}
  fmt.Println(v)
  v.ScaleFunc(2)
  fmt.Println(v)
  Scale(&v, 2)
  fmt.Println(v)

  p := &Vertex{5,6}
  fmt.Println(p)
  p.ScaleFunc(7)
  fmt.Println(p)
  Scale(p, 10)
  fmt.Println(p)

}
