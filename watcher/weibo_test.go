package watcher

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func testWeiboCard(t *testing.T, testFile string) {
	payload, err := ioutil.ReadFile(testFile)
	if err != nil {
		t.Error(err)
	}

	var m card
	err = json.Unmarshal([]byte(payload), &m)
	if err != nil {
		t.Error(err)
	}

	watcher := weiboWatcher{
		latestMblog: m.Mblog,
	}

	t.Log(watcher.parseContent())
}

func TestParseWeibo01(t *testing.T) {
	testWeiboCard(t, "tests/01-mblog-with-tag-and-pic.json")
}

func TestParseWeibo02(t *testing.T) {
	testWeiboCard(t, "tests/02-mblog-with-video.json")
}

func TestParseWeibo03(t *testing.T) {
	testWeiboCard(t, "tests/03-mblog-with-text.json")
}

func TestParseWeibo04(t *testing.T) {
	testWeiboCard(t, "tests/04-mblog-with-article.json")
}
