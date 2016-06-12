package main

import (
	"encoding/json"
	"github.com/apex/go-apex"
	"github.com/matthew-andrews/semver/latest"
	"github.com/matthew-andrews/semver/sources"
)

type message struct {
	Source string `json:"source"`
	Id     string `json:"id"`
}

func main() {
	apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
		var m message
		if err := json.Unmarshal(event, &m); err != nil {
			return nil, err
		}

		client, err := sources.New(m.Source)
		if err != nil {
			return nil, err
		}

		versions, err := client(m.Id)
		if err != nil {
			return nil, err
		}
		return latest.Latest(versions)
	})
}
