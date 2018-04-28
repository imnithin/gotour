package main

import "fmt"
import "strings"

func main() {

  tic_tac := [][]string {
    []string{"-", "-", "-"},
    []string{"-", "-", "-"},
    []string{"-", "-", "-"},
  }

  tic_tac[0][0] = "X"
  tic_tac[2][2] = "X"

  for i :=0; i < len(tic_tac); i++ {
    fmt.Printf("%s\n", strings.Join(tic_tac[i], " "))
  }
}
