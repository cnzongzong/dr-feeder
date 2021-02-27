package notifier

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/hguandl/dr-feeder/v2/common"
	"github.com/hguandl/dr-feeder/v2/notifier/wxmsgapp"
)

type WorkWechatNotifier struct {
	Client *wxmsgapp.WxAPIClient
}

type textCardPayload struct {
	Touser  string   `json:"touser"`
	Msgtype string   `json:"msgtype"`
	Agentid string   `json:"agentid"`
	News    textCard `json:"textcard"`
}

type textCard struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	PicURL      string `json:"picurl"`
}

// FromWxAPIClient creates a Notifier with an API client.
func FromWxAPIClient(client *wxmsgapp.WxAPIClient) Notifier {
	return WorkWechatNotifier{Client: client}
}

func formatText(payload common.NotifyPayload) (string, string) {
	var title, desc string

	firstParaIdx := strings.Index(payload.Body, "\n\n")

	// Only one paragraph
	if firstParaIdx == -1 {
		// Short content.
		if len(payload.Body) < 128 {
			title = payload.Body
			desc = "点击查看原文"
			// Long content.
		} else {
			title = payload.Title
			desc = payload.Body
		}
		return title, desc
	}

	// 1st paragraph is short. which can be seen as the title.
	if firstParaIdx <= 128 {
		title = payload.Body[:firstParaIdx]
		desc = payload.Body[firstParaIdx+2:]
		return title, desc
	}

	// 1st paragraph is too long. Use the general title.
	if firstParaIdx > 128 {
		title = payload.Title
		desc = payload.Body
		return title, desc
	}

	// Default results
	return payload.Title, payload.Body
}

func (notifier WorkWechatNotifier) Push(payload common.NotifyPayload) {
	title, desc := formatText(payload)

	data, err := json.Marshal(
		textCardPayload{
			Touser:  notifier.Client.ToUser,
			Msgtype: "news",
			Agentid: notifier.Client.AgentID,
			News: textCard{
				Title:       title,
				Description: desc,
				URL:         payload.URL,
				PicURL:      payload.PicURL,
			},
		},
	)
	if err != nil {
		log.Println(err)
		return
	}

	err = notifier.Client.SendMsg(data)
	if err != nil {
		log.Println(err)
	}
}
