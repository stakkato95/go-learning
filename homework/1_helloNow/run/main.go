package main

import (
	"fmt"
	"os"

	"github.com/stakkato95/hellownow/hellownow"
)

func main() {
	fmt.Println("hi!")
	hellownow.WriteTime(os.Stdout)
}
