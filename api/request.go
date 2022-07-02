package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/LAzyDev-101/stupid-server/app"
)

func PostChallenge(appApi *app.AppCaptcha, rw http.ResponseWriter, r *http.Request) {
	var params app.RequestParams

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Printf("error: %+v", err)
		return
	}
	defer r.Body.Close()
	log.Println(string(body))

	if err = json.Unmarshal(body, &params); err != nil {
		log.Printf("error: %+v", err)
		return
	}

	resp, err := appApi.ProcessRequest(params)
	if err != nil {
		log.Printf("error: %+v", err)
		return
	}
	log.Printf("resp: %v", resp)

	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(rw).Encode(resp); err != nil {
		log.Printf("error: %+v", err)
		return
	}

}
