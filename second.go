package main

import "fmt"

func main() {
	username := "Ako sa máš"
	runes := []rune(username)
    
	for _, v := range runes {
	fmt.Print(string(v), " ")
	}
	logLevel := "お元気ですか"
	for index, runeVal := range logLevel {
		fmt.Println(index, string(runeVal))
	}
}