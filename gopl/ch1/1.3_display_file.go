// On a Wednesday's night before watching a movie...
//
// A program that displays the content of a file.
// It defaults to the program's source file.
package main

import (
  "bufio"
  "fmt"
  "os"
)

func main(){
  files := os.Args[1:]
  var filename string
  if len(files) == 0{
    // open current's program source code
    filename = os.Args[0] + ".go"
  } else {
    filename = os.Args[1]
  }
  fmt.Println("Opening " + filename + "...")
  file, err := os.Open(filename)
  if err != nil {
    fmt.Println("Error!")
    fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
  }  else{
    displayFile(file)
  }
}

func displayFile(file *os.File){
  input := bufio.NewScanner(file)

  for input.Scan(){
    fmt.Println(input.Text())
  }
  file.Close()
}