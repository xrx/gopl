// On a christmas day...
package main

import (
  "fmt"
  "io/ioutil"
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
  fileData, err := ioutil.ReadFile(filename)
  if err != nil {
    fmt.Println("Error!")
    fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
  }  else{
    displayFile(fileData)
  }
}

func displayFile(fileData []byte){
  var sep string
  for _, charData := range fileData{
    if charData == 10{
      sep = "\n"
      fmt.Printf(sep)
    } else if charData == 32 {
      sep = " "
      fmt.Printf(sep)
    } else {
      fmt.Printf("%v", charData)
    }
  }
}