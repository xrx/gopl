// exercise 1.2
package main

import (
  "fmt"
  "os"
  "strconv"
)

func main() {
  s, sep := "", ""
  sep = " "
  for index, arg := range os.Args[1:] {
    index += 1
    s = strconv.Itoa(index) + sep + arg
    fmt.Println(s)
  }

}

//!-
