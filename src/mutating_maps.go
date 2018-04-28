package main
import "fmt"
func main() {
  m := make(map[string]int)
  m["Apples"] = 1 // insert or update
  fmt.Println("The Value is:", m["Apples"])
  m["Oranges"] = 2 // insert or update
  fmt.Println("The Value is:", m["Oranges"])
  m["Oranges"] = 3 // insert or update
  fmt.Println("The Value is:", m["Oranges"])
  fmt.Println("m:", m)
  delete(m, "Oranges")
  fmt.Println("m:", m)
  ele, exists := m["Apples"]
  fmt.Println("ele: ", ele, exists)
  delete(m, "Apples")
  ele1, exists1 := m["Apples"]
  fmt.Println("ele: ", ele1, exists1)
}
