package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func home(w http.ResponseWriter, req *http.Request) {
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.Write([]byte("an error occurred"))
		w.WriteHeader(http.StatusInternalServerError)
	}

	fmt.Printf("Message from server:\n%s\n", b)

	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/", home)
	fmt.Errorf("Listening to server", http.ListenAndServe(":8080", nil))
}