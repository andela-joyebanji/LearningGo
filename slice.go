package main

import "fmt"


func main() {
  //create a slice of string that contains a length and
  //capacity of 5 elements
  slice := make([]int, 5)
  slice[3] = 10
  sli := append(slice, 20)
  sli2 := append(sli, 20)

  fmt.Println(slice)
  fmt.Println(sli)
  fmt.Println(sli2)

  slice[0] = 56
  slice[1] = 89
  slice[2] = 78
  slice[3] = 65
  slice[4] = 999

  sli3 := append(slice, 9999)

  fmt.Println(sli3)

  slice4 := []int{10, 20, 30, 40}

  slice5 := append(slice4, 50)

  fmt.Println(slice5)

  names := []string{"Oyebanji", "Jacob"}

  fmt.Println(names[1:1])

}