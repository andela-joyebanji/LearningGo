package main

import (
  "fmt"
  "strings"
)

func main() {
  name := "Oyebanji Jacob Mayowa"
  fmt.Println(strings.Contains(name, "Jacob"))
  fmt.Println(strings.Index(name, "Jacob" ))
  fmt.Println(strings.Count(name, "a" ))
  fmt.Println(strings.Replace(name, "a" , "x", 1))
  fmt.Println(strings.Replace(name, "a" , "x", 3))
  split_name := strings.Split(name, "a")
  fmt.Println(split_name)
  fmt.Println(strings.Join(split_name, "^"))

}