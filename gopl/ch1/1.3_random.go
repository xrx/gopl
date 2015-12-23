package main

import(
  "fmt"
  "strconv"
)

func main(){
  key_values := make(map[string]int)
  names := make(map[int]string)
  key_values["hello"]++
  key_values["goodbye"] = 20
  names[20], names[1] = "luke", "han"

  fmt.Println("** Printing key_values **")
  for key, count := range key_values{
    fmt.Println("Count " + key + ": " + strconv.Itoa(count))
  }

  fmt.Println("** Printing names **")
  for index, name := range names{
    index++
    fmt.Println("name " + strconv.Itoa(index) + ": " + name)
  }
}