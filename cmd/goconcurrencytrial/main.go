package main

import (
	"goconcurrencytrial/internal/goconcurrencytrial"

	"strings"
)

func main() {
	str1 := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."
	list1 := strings.Split(str1, " ")

	goconcurrencytrial.DoWorkOnList(list1)
}
