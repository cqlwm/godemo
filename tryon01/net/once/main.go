package main

import (
	"log"
	"net/http"
	"tryon01/net/once/config/datasource"
	"tryon01/net/once/handler/video"
)

func main() {
	log.SetFlags(log.Llongfile | log.LstdFlags)

	http.HandleFunc("/query", video.QueryHandler)

	if err := datasource.Build(); err != nil {
		log.Println("Failed to connect to the database")
		panic(err)
	}

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println("Server startup failed")
		panic(err)
	}
}
