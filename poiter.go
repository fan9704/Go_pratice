package go_project

import "fmt"
func main2() {
	var count1  *int
	count2 := new(int)
	countTemp := 5
	count3 := &countTemp

	if count1 != nil {
		fmt.Printf("count1: %#v\n", *count1)
	}
	if count2 != nil {
		fmt.Printf("count2: %#v\n", *count2)
	}
	if count3 != nil {
		fmt.Printf("count3: %#v\n", *count3)
	}
}
