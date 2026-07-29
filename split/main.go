package main

import "fmt"

type Split struct {
	newString     string
	separedValues []string
}

func (s *Split) split(orginalString, separator string) []string {
	var result []string
	current := ""

	for i := 0; i < len(orginalString); i++ {
		if string(orginalString[i]) == separator {
			result = append(result, current)
			current = ""
		} else {
			current += string(orginalString[i])
		}
	}

	result = append(result, current)
	return result
}

func main() {
	n := Split{}

	fmt.Println(n.split("Hello_World!", "_"))
}
