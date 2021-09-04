package main

import (
	"fmt"
	"log"

	"github.com/stakkato95/unpack_string/unpack"
)

func main() {
	if result, err := unpack.Unpack(`qwe\\5a`); err != nil {
		log.Fatalln("error result for input: a4bc2d5e")
	} else {
		fmt.Println(result)
	}
}
