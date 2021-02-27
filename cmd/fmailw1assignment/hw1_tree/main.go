package main

import (
	"fmt"
	"io"
	"os"
)

const ALL_FILES = -1

func main() {
	_, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	dirTree(os.Stdout, os.Args[1], false)
}

func dirTree(writer io.Writer, path string, printFiles bool) error {
	dirTreeWithLevel(writer, path, printFiles, 0, "")
	return nil
}

func dirTreeWithLevel(writer io.Writer, path string, printFiles bool, level int, prefix string) {
	file, _ := os.Open(path)
	allFiles, _ := file.Readdir(ALL_FILES)

	for i, f := range allFiles {
		branchChar := "├"
		if i == len(allFiles)-1 {
			branchChar = "└"
		}
		fmt.Printf("%s%s───%s\n", prefix, branchChar, f.Name())

		newPrefix := "│   "
		if i == len(allFiles)-1 {
			newPrefix = "    "
		}
		if f.IsDir() {
			dirTreeWithLevel(
				writer,
				fmt.Sprintf("%s%s%s", path, string(os.PathSeparator), f.Name()),
				printFiles,
				level+1,
				fmt.Sprintf("%s%s", prefix, newPrefix))
		}
	}
}

func getLastDir(allFiles []os.FileInfo) os.FileInfo {
	length := len(allFiles)
	for i := range allFiles {
		i_reversed := length - i - 1
		if allFiles[i_reversed].IsDir() {
			return allFiles[i_reversed]
		}
	}

	return nil
}
