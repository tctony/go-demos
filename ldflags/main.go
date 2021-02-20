package main

import (
	"fmt"

	"github.com/tctony/go-demos/ldflags/some_package"
)

var name = "unknown"

func main() {
	fmt.Printf("name: %s\n", name)

	fmt.Printf("some_package.Value: %s\n", some_package.Value)
}
