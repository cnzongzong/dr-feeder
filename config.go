package main

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/hguandl/dr-feeder/v2/notifier"
	"github.com/hguandl/dr-feeder/v2/notifier/wxmsgapp"
	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v2"
)

type yamlConfig struct {
	Version   string
	Notifiers []map[string]interface{}
}

// ParseConfig loads from config file and returns a list of Notifiers.
func ParseConfig(path string) ([]notifier.Notifier, error) {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config yamlConfig
	err = yaml.Unmarshal([]byte(yamlFile), &config)
	if err != nil {
		return nil, err
	}

	if config.Version != "1.0" {
		return nil, errors.New("Invalid config version")
	}

	ret := make([]notifier.Notifier, len(config.Notifiers))
	for idx, ntfc := range config.Notifiers {
		ntft, ok := ntfc["type"].(string)
		if !ok {
			err = errors.New("Invalid notifier config")
			break
		}

		switch ntft {
		case "custom":
			var ntf notifier.CustomNotifier
			err = mapstructure.Decode(ntfc, &ntf)
			ret[idx] = ntf
		case "tgbot":
			var ntf notifier.TgBotNotifier
			err = mapstructure.Decode(ntfc, &ntf)
			ret[idx] = ntf
		case "workwx":
			var wxClient wxmsgapp.WxAPIClient
			err = mapstructure.Decode(ntfc, &wxClient)
			ret[idx] = notifier.WorkWechatNotifier{Client: &wxClient}
		case "bark":
			var ntf notifier.BarkNotifier
			err = mapstructure.Decode(ntfc, &ntf)
			ret[idx] = ntf
		case "ifttt":
			var ntf notifier.IFTTTNotifier
			err = mapstructure.Decode(ntfc, &ntf)
			ret[idx] = ntf
		default:
			err = fmt.Errorf("Unknown notifier type \"%s\"", ntft)
		}

		if err != nil {
			break
		}
	}

	return ret, err
}
