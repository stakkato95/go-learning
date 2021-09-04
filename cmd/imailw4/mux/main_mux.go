package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	go runServer(":8081")
	go runServer(":8082")
	runServerForStatics()
}

func runServer(port string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	server := http.Server{
		Addr:         port,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("starting server at port ", port)
	server.ListenAndServe()
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hi!")
}

func runServerForStatics() {
	mux := http.NewServeMux()

	staicPrefix := "/data/"
	mux.Handle(staicPrefix, http.StripPrefix(staicPrefix, http.FileServer(http.Dir("./static"))))

	mux.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		fmt.Println(string(body))
		//unmarshalling
	})

	server := http.Server{
		Addr:         ":8083",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("starting server at port 8083")
	server.ListenAndServe()
}
