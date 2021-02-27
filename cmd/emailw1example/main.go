package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	if err := uniqFromUnsortedInput(os.Stdin, os.Stdout); err != nil {
		panic(err.Error())
	}
}

func uniqFromUnsortedInput(input io.Reader, output io.Writer) error {
	//pass data to stdin
	//cat cmd/dmailw1example/nums.txt | go run cmd/dmailw1example/main.go
	in := bufio.NewScanner(input)
	alreadySeen := make(map[string]bool)

	for in.Scan() {
		txt := in.Text()

		if unicode.IsLetter(rune(txt[0])) {
			return fmt.Errorf("Letter found in input, but only number are expected")
		}

		if _, isValueAlreadySeen := alreadySeen[txt]; !isValueAlreadySeen {
			alreadySeen[txt] = true
			fmt.Fprintln(output, txt)
		}
	}

	return nil
}
