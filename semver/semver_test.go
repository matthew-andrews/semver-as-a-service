package semver

import (
	"github.com/matthew-andrews/go-version"
	"testing"
)

type mockSource struct{}

func getVersion(versionString string) *version.Version {
	v, _ := version.NewVersion(versionString)
	return v
}

func (source mockSource) VersionsFor(id string) ([]*version.Version, error) {
	return []*version.Version{
		getVersion("v0.0.1"),
		getVersion("v0.1.1"),
		getVersion("v1.1.1"),
		getVersion("v1.0.3"),
		getVersion("v1.0.2"),
		getVersion("v0.1.2"),
	}, nil
}

func TestLatest(t *testing.T) {
	source := mockSource{}

	actual, err := Semver(source, "id", "latest")

	if err != nil {
		t.Fatalf("unexpected error thrown", err)
	}

	expected := "v1.1.1"

	if actual != expected {
		t.Fatalf("latest expected: %s, actual: %s", expected, actual)
	}
}

func TestTilde(t *testing.T) {
	source := mockSource{}

	actual, err := Semver(source, "id", "~1.0.0")

	if err != nil {
		t.Fatalf("unexpected error thrown", err)
	}

	expected := "v1.0.3"

	if actual != expected {
		t.Fatalf("latest expected: %s, actual: %s", expected, actual)
	}
}
