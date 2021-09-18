package main

import (
	"flag"

	"github.com/stakkato95/copyutil/internal/copyutil"
)

var (
	chunkSize    int64
	copyBytes    int64
	offset       int64
	showProgress bool
	from         string
	to           string
)

func init() {
	flag.Int64Var(&chunkSize, "bs", 1, "chuck of bytes to copy within one iteration (default 1)")
	flag.Int64Var(&copyBytes, "limit", 1, "limit of bytes to copy")
	flag.Int64Var(&offset, "offset", 1, "offset in input file")
	flag.BoolVar(&showProgress, "progress", true, "show progress in stdout")
	flag.StringVar(&from, "from", "", "file to read from (mandatory)")
	flag.StringVar(&to, "to", "", "file to write to (mandatory)")
}

func main() {
	copyutil.Copy(copyutil.CopyConfig{
		ChunkSize:    chunkSize,
		CopyBytes:    copyBytes,
		Offset:       offset,
		ShowProgress: showProgress,
		From:         from,
		To:           to,
	})
}
