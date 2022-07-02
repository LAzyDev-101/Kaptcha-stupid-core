package main

import (
	"log"
	"net/http"

	"github.com/LAzyDev-101/stupid-server/api"
	"github.com/LAzyDev-101/stupid-server/app"
	"github.com/gorilla/mux"
)

func serveStatic(path string) http.Handler {
	return http.StripPrefix("static", http.FileServer(http.Dir(path)))
}

func main() {

	app := &app.AppCaptcha{
		Users: make(map[string][]string),
	}
	router := mux.NewRouter()

	// router.Handle(
	// 	"/memory-game",
	// 	http.StripPrefix("/memory-game/",
	// 		serveStatic("static/memory-game")),
	// )

	// router.Handle("/memory-game", http.StripPrefix("/static/memory-game", fs))
	s := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
	router.PathPrefix("/static/").Handler(s)

	router.HandleFunc(
		"/post_finish",
		func(w http.ResponseWriter, r *http.Request) {
			api.PostChallenge(app, w, r)
		},
	)

	log.Print("Listening on :3000...")
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		log.Fatal(err)
	}
}
