package main

import (
	"bufio"
	"fmt"
	"os"
	s "strings"
)

// Read standard in text until a newline Python style
func Input(prefix string) (input string) {

	fmt.Print(prefix)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	return s.TrimRight(text, "\r\n")
}
