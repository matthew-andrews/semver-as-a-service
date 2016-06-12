package sources

import (
	"errors"
	"fmt"
	"github.com/matthew-andrews/go-version"
)

type Source interface {
	VersionsFor(string) ([]*version.Version, error)
}

func New(source string) (Source, error) {
	if source == "github" {
		return &githubSource{}, nil
	}
	return nil, errors.New(fmt.Sprintf("unsupported source %s", source))
}
