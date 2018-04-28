package main
import (
  "fmt"
  "time"
)

func main() {
  t := time.Now();
  switch {
    case t.Hour() < 12: fmt.Println("Gm")
    case t.Hour() < 4: fmt.Println("Ge")
    // case t.Hour() > 4: fmt.Println("Gn")
    default: fmt.Println("Gn")
  }
}
