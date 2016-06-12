package main

import (
	"encoding/json"
	"github.com/apex/go-apex"
	"github.com/matthew-andrews/semver/semver"
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

		return semver.Semver(m.Source, m.Id, m.Satisfies)
	})
}
