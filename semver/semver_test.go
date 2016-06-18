package semver

import (
	"errors"
	"github.com/matthew-andrews/go-version"
	"testing"
)

type mockSource struct{}

func getVersion(versionString string) *version.Version {
	v, _ := version.NewVersion(versionString)
	return v
}

func (source mockSource) VersionsFor(id string) ([]*version.Version, error) {
	if id == "erroring" {
		return []*version.Version{}, errors.New("VersionsFor errored for some reason")
	}
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

func TestErrorNoVersionsMatchingConstraint(t *testing.T) {
	source := mockSource{}
	_, err := Semver(source, "id", "~2.0.0")

	if err == nil {
		t.Fatalf("error expected to be thrown but none was")
	}
	expected := "no version found matching constraint"
	if err.Error() != expected {
		t.Fatalf("latest expected: %s, actual: %s", expected, err.Error())
	}
}

func TestErrorBadConstraint(t *testing.T) {
	source := mockSource{}
	_, err := Semver(source, "id", "BAD CONSTRAINT")

	if err == nil {
		t.Fatalf("error expected to be thrown but none was")
	}
	expected := "Malformed constraint: BAD CONSTRAINT"
	if err.Error() != expected {
		t.Fatalf("latest expected: %s, actual: %s", expected, err.Error())
	}
}

func TestErrorVersionsForErrored(t *testing.T) {
	source := mockSource{}
	_, err := Semver(source, "erroring", "^1.0.0")

	if err == nil {
		t.Fatalf("error expected to be thrown but none was")
	}
	expected := "VersionsFor errored for some reason"
	if err.Error() != expected {
		t.Fatalf("latest expected: %s, actual: %s", expected, err.Error())
	}
}
