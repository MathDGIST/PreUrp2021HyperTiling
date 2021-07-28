package main 

import (
  "fmt"
  "main/matrix"
)

func main() {
  fmt.Println("Hello, World!")
  mb := matrix.Matrix{[][]complex128{{1, 2}, {3, 4}}}
  fmt.Println(mb)
}