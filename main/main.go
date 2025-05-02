package main

import (
	"fmt"
	"log"
	"net/http"

	htmlGenerator "github.com/sohWenMing/cyoa/html_generator"
	jsonDecoder "github.com/sohWenMing/cyoa/json_decoder"
	"github.com/sohWenMing/cyoa/server"
)

func main() {
	storyChoices, err := jsonDecoder.GetOptionsFromJSON()
	if err != nil {
		log.Fatal("error occured when initialising server: ", err)
	}
	introHTML, err := htmlGenerator.GenerateHTML(storyChoices["intro"])
	if err != nil {
		log.Fatal("error occured when initialising server: ", err)
	}
	mux := server.InitServer()
	mux.Handle("/", server.DefaultHandler(introHTML))
	fmt.Println("listening on port 8080")
	http.ListenAndServe(":8080", mux)
}
