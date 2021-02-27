package notifier

import (
	"log"
	"net/http"
	"net/url"

	"github.com/hguandl/dr-feeder/v2/common"
)

type CustomNotifier struct {
	APIURL string `mapstructure:"api_url"`
}

func (notifier CustomNotifier) Push(payload common.NotifyPayload) {
	r, err := http.PostForm(notifier.APIURL,
		url.Values{
			"title":  {payload.Title},
			"body":   {payload.Body},
			"url":    {payload.URL},
			"picurl": {payload.PicURL},
		})
	if err != nil {
		log.Println(err)
		return
	}
	defer r.Body.Close()
}
