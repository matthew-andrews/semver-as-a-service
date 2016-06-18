package semver

import (
	"errors"
	"github.com/matthew-andrews/go-version"
	"sort"
)

type sourceInterface interface {
	VersionsFor(string) ([]*version.Version, error)
}

func Semver(source sourceInterface, id string, satisfies string) (string, error) {
	versions, err := source.VersionsFor(id)
	if err != nil {
		return "", err
	}

	sort.Reverse(version.Collection(versions))
	if satisfies == "latest" {
		return versions[0].Raw(), nil
	}

	constraint, err := version.NewConstraint(satisfies)
	if err != nil {
		return "", err
	}

	for _, version := range versions {
		if constraint.Check(version) {
			return version.Raw(), nil
		}
	}

	return "", errors.New("no version found matching constraint")
}
