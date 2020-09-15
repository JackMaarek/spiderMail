package main

import (
	"fmt"
	"net/http"
	//"github.com/JackMaarek/spiderMail/consummer"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8088", nil)
	//consummer.ReceiveToRabbit()
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
