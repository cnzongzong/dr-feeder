package notifier

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hguandl/dr-feeder/v2/common"
)

type BarkNotifier struct {
	Tokens []string
}

func (notifier BarkNotifier) apiURL() string {
	return "https://api.day.app"
}

func (notifier BarkNotifier) Push(payload common.NotifyPayload) {
	for _, token := range notifier.Tokens {
		r, err := http.Get(fmt.Sprintf(
			"%s/%s/%s/%s?url=%s",
			notifier.apiURL(),
			token,
			payload.Title,
			payload.Body,
			payload.URL,
		))
		if err != nil {
			log.Println(err)
		} else {
			r.Body.Close()
		}
	}
}
