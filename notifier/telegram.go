package notifier

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/hguandl/arknights-news-watcher/v2/common"
)

type tgBotNotifer struct {
	botAPIKey string
	chats     []string
}

func NewTgBotNotifier(botAPIKey string, chats []string) Notifier {
	return tgBotNotifer{
		botAPIKey: botAPIKey,
		chats:     chats,
	}
}

func FromTgBotNotifierConfig(config map[string]interface{}) (Notifier, bool) {
	botAPIKey, ok := config["api_key"].(string)
	if !ok {
		return nil, false
	}

	chatsRaw, ok := config["chats"].([]interface{})
	if !ok {
		return nil, false
	}

	chats := make([]string, len(chatsRaw))
	for i := range chats {
		chats[i], ok = chatsRaw[i].(string)
		if !ok {
			return nil, false
		}
	}

	return NewTgBotNotifier(botAPIKey, chats), true
}

func (notifier tgBotNotifer) apiURL() string {
	return fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage",
		notifier.botAPIKey,
	)
}

func (notifer tgBotNotifer) Push(payload common.NotifyPayload) {
	texts := payload.Body + "\n\n" + payload.URL

	for _, chat := range notifer.chats {
		_, err := http.PostForm(notifer.apiURL(),
			url.Values{
				"chat_id": {fmt.Sprint(chat)},
				"text":    {texts},
			})
		if err != nil {
			log.Println(err)
		}
	}

}
