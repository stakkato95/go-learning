package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)

	http.HandleFunc("/pages/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "pages")
		fmt.Fprintln(w, r.URL.String())
	})

	msgSrv := &MesagingService{Token: "100500token"}
	http.Handle("/message", msgSrv)

	fmt.Println("starting server at :8081")
	http.ListenAndServe(":8081", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
	w.Write([]byte("!!!!!!!"))
}

//service + DI
type MesagingService struct {
	Token string
}

func (this *MesagingService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "handled by a service with token %s\n", this.Token)
}
