package main


import (
  "fmt"
  "bufio"
  "os"
  "strings"
)

func main() {
  var age int
  var name string
  var err error
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("Enter your name: ")
  if name, err = reader.ReadString('\n'); err != nil {
    fmt.Println("Error")
    return
  }
  fmt.Print("Enter your age: ")
  fmt.Scanf("%d", &age)
  name = strings.TrimRight(name, "\n")
  if age >= 18 {
    fmt.Printf("You can vote %s!!!", name)
  } else {
    fmt.Printf("Sorry %s,you can't vote!!!", name)
  }
}