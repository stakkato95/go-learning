package hellonow

import (
	"io"
	"log"
	"time"

	"github.com/beevik/ntp"
)

const host = "0.at.pool.ntp.org"

func PrintTime(in io.Writer) {
	log.SetOutput(in)

	if time, err := ntp.Time(host); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("current time: %s\n", time.String())
	}

	if response, err := ntp.Query(host); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("exact time: %s\n", time.Now().Add(response.ClockOffset))
	}

	log.Println("time ist written correctly")
}
