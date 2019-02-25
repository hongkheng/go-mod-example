package main

import (
	"fmt"

	"github.com/hongkheng/gohello"
)

func main() {
	personName := gohello.Hello("hk")
	fmt.Println("Test print %s", personName)
}
