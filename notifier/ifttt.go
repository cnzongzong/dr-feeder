package notifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/hguandl/dr-feeder/v2/common"
)

type IFTTTNotifier struct {
	Webhooks []webhookConfig
}

type webhookConfig struct {
	Event  string
	APIKey string `mapstructure:"api_key"`
}

type webhookPayload struct {
	Value1 string `json:"value1"`
	Value2 string `json:"value2"`
	Value3 string `json:"value3"`
}

func (notifier IFTTTNotifier) apiURL(webhook webhookConfig) string {
	return fmt.Sprintf("https://make.ifttt.com/trigger/%s/with/key/%s",
		webhook.Event,
		webhook.APIKey)
}

func (notifier IFTTTNotifier) Push(payload common.NotifyPayload) {
	for _, webhook := range notifier.Webhooks {
		data, err := json.Marshal(
			webhookPayload{
				Value1: payload.Body,
				Value2: payload.URL,
				Value3: payload.PicURL,
			},
		)
		if err != nil {
			log.Println(err)
			return
		}

		r, err := http.Post(
			notifier.apiURL(webhook),
			"application/json",
			bytes.NewBuffer(data),
		)
		if err != nil {
			log.Println(err)
		} else {
			r.Body.Close()
		}
	}
}
