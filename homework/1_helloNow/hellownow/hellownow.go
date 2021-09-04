package hellownow

import (
	"fmt"
	"io"
	"time"
)

func WriteTime(in io.Writer) {
	y, m, d := time.Now().Date()
	fmt.Fprintf(in, "%d-%d-%d", d, m, y)
}
