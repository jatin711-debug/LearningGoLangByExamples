package main

import (
	"errors"
	"fmt"
)

func divError() (int, error) {
	return 7, errors.New("Packagr error")
}

func main() {
	val, err := divError()
	fmt.Println("val:", val, " err:", err)
}
