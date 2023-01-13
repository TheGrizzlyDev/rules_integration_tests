package main

import (
	"fmt"
	"flag"
)

var (
	port = flag.Int("port", 12345, "")
)

func main() {
	flag.Parse()
	fmt.Println("Hola")
}