package util

import (
	"fmt"
	"strings"
)

func Section(name string) {
	delim := strings.Repeat("=", len(name))
	fmt.Println()
	fmt.Println(name)
	fmt.Println(delim)
}

func Spacer() {
	fmt.Println("--------")
}
