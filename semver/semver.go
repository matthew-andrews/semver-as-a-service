package semver

import (
	"errors"
	"github.com/matthew-andrews/semver/latest"
	"github.com/matthew-andrews/semver/sources"
)

func Semver(sourceName string, id string, satisfies string) (string, error) {
	if satisfies != "latest" {
		return "", errors.New("satisfies must equal latest")
	}

	source, err := sources.New(sourceName)
	if err != nil {
		return "", err
	}

	versions, err := source.VersionsFor(id)
	if err != nil {
		return "", err
	}

	return latest.Latest(versions), nil
}
