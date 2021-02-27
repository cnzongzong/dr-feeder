package watcher

import "time"

type indexData struct {
	Data struct {
		UserInfo struct {
			ScreenName string `json:"screen_name"`
		} `json:"userInfo"`
		TabsInfo struct {
			Tabs []struct {
				TabType     string `json:"tab_type"`
				Containerid string `json:"containerid"`
			} `json:"tabs"`
		} `json:"tabsInfo"`
	} `json:"data"`
}

type mblog struct {
	CreatedAt string                 `json:"created_at"`
	ID        string                 `json:"id"`
	Text      string                 `json:"text"`
	PicURL    string                 `json:"original_pic,omitempty"`
	PageInfo  map[string]interface{} `json:"page_info"`
}

type card struct {
	CardType int   `json:"card_type"`
	Mblog    mblog `json:"mblog,omitempty"`
}

type cardData struct {
	Data struct {
		Cards []card `json:"cards"`
	} `json:"data"`
}

type weiboWatcher struct {
	uid         uint64
	updateTime  time.Time
	containerID string
	name        string
	latestMblog mblog
}

type pageInfo struct {
	Type    string
	PagePic struct {
		URL string
	} `mapstructure:"page_pic"`
}
