package notifier

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/hguandl/dr-feeder/v2/common"
)

type tgBotNotifier struct {
	botAPIKey string
	chats     []string
}

// NewTgBotNotifier creates a Notifier of Telegram Bot.
func NewTgBotNotifier(botAPIKey string, chats []string) Notifier {
	return tgBotNotifier{
		botAPIKey: botAPIKey,
		chats:     chats,
	}
}

// FromTgBotNotifierConfig parses the config to create a tgBotNotifier.
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

func (notifier tgBotNotifier) apiURL() string {
	return fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage",
		notifier.botAPIKey,
	)
}

func (notifier tgBotNotifier) Push(payload common.NotifyPayload) {
	texts := payload.Body + "\n\n" + payload.URL

	for _, chat := range notifier.chats {
		_, err := http.PostForm(notifier.apiURL(),
			url.Values{
				"chat_id": {fmt.Sprint(chat)},
				"text":    {texts},
			})
		if err != nil {
			log.Println(err)
		}
	}
}
