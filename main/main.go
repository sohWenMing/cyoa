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
	if err != nil {
		log.Fatal("error occured when initialising server: ", err)
	}
	storyHTMLS, err := htmlGenerator.GetHTMLS(storyChoices)
	if err != nil {
		log.Fatal("error occured when initialising server: ", err)
	}

	mux := server.InitServer()
	mux.Handle("/", server.DefaultHandler(storyHTMLS["intro"]))
	mux.Handle("/new-york", server.DefaultHandler(storyHTMLS["new-york"]))
	mux.Handle("/debate", server.DefaultHandler(storyHTMLS["debate"]))
	mux.Handle("/sean-kelly", server.DefaultHandler(storyHTMLS["sean-kelly"]))
	mux.Handle("/mark-bates", server.DefaultHandler(storyHTMLS["mark-bates"]))
	mux.Handle("/denver", server.DefaultHandler(storyHTMLS["denver"]))
	mux.Handle("/home", server.DefaultHandler(storyHTMLS["home"]))
	fmt.Println("listening on port 8080")
	http.ListenAndServe(":8080", mux)
}
