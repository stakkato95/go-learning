package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"time"
)

func main() {
	go runServer()
	time.Sleep(1 * time.Second)

	sendClientRequest()
}

func runServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintln(w, "error when reading body")
			return
		}

		defer r.Body.Close()
		fmt.Fprintln(w, "you sent", string(body))
	})

	server := http.Server{
		Addr:         ":8081",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	server.ListenAndServe()
}

func sendClientRequest() {
	//create client
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	client := &http.Client{
		Timeout:   10 * time.Second,
		Transport: transport,
	}

	//send request
	data := `{"id": 100500, "name": "alex"}`
	body := bytes.NewBufferString(data)
	url := "http://127.0.0.1:8081/json"
	req, _ := http.NewRequest(http.MethodPost, url, body)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.Itoa(len(data)))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error happend:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("no errors")

	if body, err := ioutil.ReadAll(resp.Body); err == nil {
		fmt.Println(string(body))
		return
	}

	fmt.Println("error when reading body")
}
