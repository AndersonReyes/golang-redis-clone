package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func HandlePingCommand(request string) string {
	return "PONG"
}

func HandleConnection(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	command := string(body)
	var res string

	if strings.HasPrefix(command, "PING") {
		res = HandlePingCommand(command)
	} else {
		res = "error unrecognized command"
	}

	fmt.Fprintf(w, "%s", res)
}

func main() {
	fmt.Println("Hello, world.")
	mux := http.NewServeMux()

	mux.HandleFunc("/", HandleConnection)

	err := http.ListenAndServe(":7777", mux)
	log.Fatal(err)
}
