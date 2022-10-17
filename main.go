package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	s := os.Args[1:]
	str := ""
	for _, v := range s {
		str += v
	}
	count := utf8.RuneCountInString(str)
	fmt.Println(count)
	fmt.Println(str)
	t := isContact(&str)
	fmt.Println(t)
}
func isContact(s *string) bool {
	strs := []string{
		"contact",
		"联系",
		"联络",
		"聯繫",
		"聯絡",
	}
	for _, subStr := range strs {
		if contains(s, subStr) {
			return true
		}
	}
	return false
}

func contains(s *string, subStr string) bool {
	str := strings.ReplaceAll(*s, " ", "")
	str = strings.ToLower(str)
	return strings.Contains(str, subStr)
}
