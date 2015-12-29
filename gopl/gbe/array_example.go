package main

import "fmt"

var arr [5]int

func main(){
  // arrays are zero-valued
  for _, elem := range arr{
    fmt.Println(elem)
  }

  b := [5]int{1, 2, 3, 4, 5}

  for _, elem := range b{
    fmt.Println(elem)
  }
}