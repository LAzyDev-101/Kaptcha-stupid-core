package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/LAzyDev-101/stupid-server/api"
	"github.com/LAzyDev-101/stupid-server/app"
	"github.com/gorilla/mux"
)

func serveStatic(path string) http.Handler {
	return http.StripPrefix("static", http.FileServer(http.Dir(path)))
}

func accessControlMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}
		body, _ := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		log.Printf("%s %s %s", r.Method, r.URL, string(body))
		next.ServeHTTP(w, r)
	})
}

func main() {

	app := &app.AppCaptcha{
		Users: make(map[string][]string),
	}
	router := mux.NewRouter()
	router.Use(accessControlMiddleware)

	// router.Handle(
	// 	"/memory-game",
	// 	http.StripPrefix("/memory-game/",
	// 		serveStatic("static/memory-game")),
	// )

	// router.Handle("/memory-game", http.StripPrefix("/static/memory-game", fs))
	s := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
	router.PathPrefix("/").Handler(s)

	router.HandleFunc(
		"/post_finish",
		func(w http.ResponseWriter, r *http.Request) {
			api.PostChallenge(app, w, r)
		},
	).Methods("POST")

	log.Print("Listening on :3000...")
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		log.Fatal(err)
	}
}
