package main

import (
	"fmt"

	"github.com/nicolascardich/practicando-paquetes-go/greet"

	"rsc.io/quote/v3"
)

func main() {
	fmt.Println("Hello")             // Hello
	fmt.Println(greet.Italian())     // Ciao Loki!
	fmt.Println(quote.HelloV3())     // Hello, world.
	fmt.Println(quote.Concurrency()) // Concurrency is not parallelism.
}
