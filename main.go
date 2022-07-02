package main

import (
	"log"
	"net/http"

	"github.com/LAzyDev-101/stupid-server/api"
)

func serveStatic(path string) http.Handler {
	return http.FileServer(http.Dir(path))
}

func main() {

	http.Handle(
		"/memory-game/",
		http.StripPrefix("/memory-game/",
			serveStatic("static/memory-game")),
	)

	http.HandleFunc("/post_finish", api.PostChallenge)

	log.Print("Listening on :3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
