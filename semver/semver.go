package semver

import (
	"errors"
	"github.com/matthew-andrews/go-version"
	"github.com/matthew-andrews/semver/sources"
	"sort"
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

	sort.Sort(version.Collection(versions))
	return versions[len(versions)-1].Raw(), nil
}
