package go_project

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)
var helloList = []string{
	"Hello, world",
	"Καλημέρα κόσμε",
	"こんにちは世界",
	" ایند مالس",
	"Привет, мир",}
func main() {
    fmt.Println("Hello Golang")
	
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(helloList))
	msg, err := hello(index)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg)
	query, limit, offset := "bat", 10, 0
	query, maxLength, offset := "apple", limit, 20
	fmt.Println("query, maxLength, offset: ", query, maxLength, offset)
	
}
func hello(index int) (string, error) {
	if index < 0 || index > len(helloList)-1 {
		return "", errors.New("out of range: " + strconv.Itoa(index))
	}
	return helloList[index], nil
}