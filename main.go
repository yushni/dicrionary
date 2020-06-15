package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloServer)
	_ = http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprint(w, "Bye bye CD, Hello CI only for master branch!")
}
