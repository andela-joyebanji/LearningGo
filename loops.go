package main

import "fmt"

func main() {
	i := 1
	for i <= 10 {
		fmt.Println("Jacob", "is", "a good", "Person 1")
		i++
	}

  //Slice
  slice := []int{10, 9, 8, 7,6,5,4,3,2,1}
  for index, value := range slice {
    fmt.Printf("Index: %d Value: %d\n", index, value)
  }
  fmt.Println("Traditional")
  for index := 2; index < len(slice); index ++ {
    fmt.Printf("Index: %d Value: %d\n", index, slice[index])
  }

  sum := 0

  for i := 1; i <= 10; i++ {
    sum += 1
  }

  fmt.Println(sum)
}
