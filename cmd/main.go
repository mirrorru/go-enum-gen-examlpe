package main

import (
	"fmt"

	"go-enum-gen-example/internal/enums"
)

func main() {
	for _, cs := range enums.ClientStatusValues() {
		fmt.Println(cs, cs.String(), cs.HumanString())
	}
}
