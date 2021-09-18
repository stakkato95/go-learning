package copyutil

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const ErrorMessageHeader = "command line arguments errors:"

type CopyConfig struct {
	ChunkSize    int64
	CopyBytes    int64
	Offset       int64
	ShowProgress bool
	From         string
	To           string
}

func Copy(config CopyConfig) {
	if errs := ValidateConfig(config); len(errs) != 0 {
		PrintErrors(errs, bufio.NewWriter(os.Stdout))
		return
	}

	CopyInternal(config)
}

func PrintErrors(errs []error, w *bufio.Writer) {
	defer w.Flush()
	w.WriteString(ErrorMessageHeader)

	for i, err := range errs {
		w.WriteString(fmt.Sprintf("\n%d. %s", i+1, err.Error()))
	}
}

func CopyInternal(config CopyConfig) error {
	from, err := os.Open(config.From)
	if err != nil {
		return fmt.Errorf("error when opening 'from' file: %w", err)
	}
	defer from.Close()

	to, err := os.Create(config.To)
	if err != nil {
		return fmt.Errorf("error when creating 'to' file: %w", err)
	}
	defer to.Close()

	_, err = from.Seek(config.Offset, io.SeekStart)
	if err != nil {
		return fmt.Errorf("could not seek to offset=%d in 'from' file: %w", config.Offset, err)
	}

	buffer := make([]byte, config.ChunkSize)
	var totalBytesCopied int64
	for totalBytesCopied < config.CopyBytes {
		_, err := from.Read(buffer)
		if err != nil {
			return fmt.Errorf("could not read bytes from 'from' file: %w", err)
		}

		_, err = to.Write(buffer)
		if err != nil {
			return fmt.Errorf("could not erite bytes to 'to' file: %w", err)
		}

		totalBytesCopied += int64(len(buffer))
		residual := config.CopyBytes - totalBytesCopied
		if residual == 0 {
			break
		}
		if residual < config.ChunkSize {
			buffer = make([]byte, residual)
		}
	}

	return nil
}
