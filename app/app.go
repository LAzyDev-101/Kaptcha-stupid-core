package app

var (
	users map[string]int
)

type RequestParams struct {
	Key string `json:"key"`
}

type RequestResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func ProcessRequest(params RequestParams) (*RequestResponse, error) {
	count, exist := users[params.Key]
	if !exist {
		users[params.Key] = 1
	} else {
		if count == 3 {
			return &RequestResponse{
				Message: "Finish",
				Status:  "FINISH",
			}, nil
		}
		users[params.Key] += 1
	}

	return &RequestResponse{
		Message: "next challenge",
		Status:  "Processing",
	}, nil
}

func init() {
	users = make(map[string]int)
}
