package go_project

import "fmt"

var level = "pkg"

func main4() {
	fmt.Println("Main start  :", level) //main()層級
	level := 42
	if true {
		fmt.Println("Block start :", level) //底下if層級
		funcA()
	}
	fmt.Println("Main end    :", level)
}

func funcA() {
	fmt.Println("funcA start :", level) //funcA()層級
}
