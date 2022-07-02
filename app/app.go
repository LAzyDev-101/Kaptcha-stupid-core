package app

import (
	"math/rand"
	"sync"
	"time"
)

type AppCaptcha struct {
	Users map[string][]string
	sync.Mutex
}

var allChallenge = []string{
	"stupid-run",
	"stupid-block",
	"stupid-bird",
	"stupid-memory",
}

type RequestParams struct {
	Key           string `json:"key"`
	ChallengeName string `json:"challenge_name"`
}

type RequestResponse struct {
	Message           string `json:"message"`
	Status            string `json:"status"`
	NextChallengeName string `json:"next_challenge_name"`
}

func (app *AppCaptcha) ProcessRequest(params RequestParams) (*RequestResponse, error) {
	app.Lock()
	defer app.Unlock()

	challengePlayed, exist := app.Users[params.Key]
	if !exist {
		app.Users[params.Key] = make([]string, 0)
	}
	if len(challengePlayed) == 4 {
		return &RequestResponse{
			Message:           "Finish",
			Status:            "FINISH",
			NextChallengeName: "",
		}, nil
	}
	app.Users[params.Key] = append(app.Users[params.Key], params.ChallengeName)

	encountered := make(map[string]bool)
	for _, playedChallenge := range app.Users[params.Key] {
		encountered[playedChallenge] = true
	}

	nextChallengeName := ""
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(allChallenge), func(i, j int) { allChallenge[i], allChallenge[j] = allChallenge[j], allChallenge[i] })

	for _, challenge := range allChallenge {
		if _, exist := encountered[challenge]; !exist {
			nextChallengeName = challenge
			break
		}
	}

	return &RequestResponse{
		Message:           "next challenge",
		Status:            "PROCESSING",
		NextChallengeName: nextChallengeName,
	}, nil
}
