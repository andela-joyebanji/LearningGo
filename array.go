package main

import (
	"fmt"
)

func main() {
	var array [5]int
	array2 := [5]int{10, 20, 30, 40, 50}
	names := [...]string{"Jacob", "Taju"}
	specificNames := [5]string{4: "Jacob", 3: "Bernard", 2: "Pius", 1: "Francis", 0: "Cecilia"}
	// Initialize index 0 and 1 of the array with integer pointers.
	array3 := [5]*int{0: new(int), 1: new(int)}

	var a int
	a = 90
	fmt.Println(array)
	fmt.Println(array2)
	fmt.Println(a)
	fmt.Println(names)
	fmt.Println(specificNames)

	array2[4] = 100
	fmt.Println(array2)
	fmt.Println(array3)

	 array4 := [2][2]int{{419, 30}, {18, 89}}
	fmt.Println(array4)
}
