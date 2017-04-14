package main

import "fmt"

func main() {
  var a, b int32 = 5, 5
  fmt.Println(Add(a,b))
}

func Add(a, b int32) int32 {
  return a + b
}
