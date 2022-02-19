package go_project

import ("fmt")

func add5Value(count int) {
	count += 5
	fmt.Println("add5Value     :", count)
}

func add5Point(count *int) {
	*count += 5
	fmt.Println("add5Point     :", *count)
}
func main3() {
	var count int
	add5Value(count)
	fmt.Println("add5Value post:", count)
	add5Point(&count)
	fmt.Println("add5Point post:", count)
}
