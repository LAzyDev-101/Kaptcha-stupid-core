package app

var (
	users map[string][]string
)

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

func ProcessRequest(params RequestParams) (*RequestResponse, error) {
	challengePlayed, exist := users[params.Key]
	if !exist {
		users[params.Key] = make([]string, 0)
	}
	if len(challengePlayed) == 4 {
		return &RequestResponse{
			Message:           "Finish",
			Status:            "FINISH",
			NextChallengeName: "",
		}, nil
	}
	users[params.Key] = append(users[params.Key], params.ChallengeName)

	encountered := make(map[string]bool)
	for _, playedChallenge := range users[params.Key] {
		encountered[playedChallenge] = true
	}

	nextChallengeName := ""
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

func init() {
	users = make(map[string][]string)
}
