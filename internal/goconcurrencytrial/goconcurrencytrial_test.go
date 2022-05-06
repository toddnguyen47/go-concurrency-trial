package goconcurrencytrial

import (
	"strings"
	"testing"
)

func Test_GivenList_When_Then(t *testing.T) {
	str1 := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."
	list1 := strings.Split(str1, " ")
	DoWorkOnList(list1)
}
