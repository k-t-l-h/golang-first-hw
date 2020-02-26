package main

import (
	"calculator/internal/pkg/calc"
	"fmt"
	"os"
)

func main() {

	result, err := calc.Calc(os.Stdin)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("result:  ", result)

}
