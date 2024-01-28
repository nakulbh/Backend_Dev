package main

import (
	"context"
	"log"
	"net/http"
	"os"
)

type TextHandler string
type TextHandlerSecond string

func main() {
	server := http.Server{
		Addr:    ":8080",
		Handler: TextHandlerSecond("hii its Nakul"),
	}
	go server.ListenAndServe()

	req, _ := http.NewRequestWithContext(context.TODO(), "GET", "http://localhost:8080", nil)
	resp, err := new(http.Client).Do(req)
	if err != nil {
		log.Fatal("err in requesting..")
	}

	defer resp.Body.Close()
	resp.Write(os.Stdout)
}

func (t TextHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write(([]byte(t)))
}

func (t TextHandlerSecond) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(t))
}
