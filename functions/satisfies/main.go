package main

import (
	"encoding/json"
	"github.com/apex/go-apex"
	"github.com/matthew-andrews/semver-as-a-service/semver"
	"github.com/matthew-andrews/semver-as-a-service/sources"
)

type message struct {
	Satisfies string `json:"satisfies"`
	Source    string `json:"source"`
	Id        string `json:"id"`
}

func main() {
	apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
		var m message
		if err := json.Unmarshal(event, &m); err != nil {
			return nil, err
		}

		source, err := sources.New(m.Source)
		if err != nil {
			return nil, err
		}
		return semver.Semver(source, m.Id, m.Satisfies)
	})
}
