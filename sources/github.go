package sources

import (
	"errors"
	"fmt"
	"github.com/google/go-github/github"
	"github.com/matthew-andrews/go-version"
	"golang.org/x/oauth2"
	"os"
	"strings"
)

type githubSource struct{}

func (g githubSource) VersionsFor(identifier string) ([]*version.Version, error) {
	var allVersions []*version.Version

	// Derive GitHub owner and repo
	parts := strings.Split(identifier, "/")
	owner, repo := parts[0], parts[1]

	// Create GitHub API client
	apiKey := os.Getenv("GITHUB_API_KEY")
	if apiKey == "" {
		return allVersions, errors.New("GITHUB_API_KEY not provided")
	}
	client := github.NewClient(oauth2.NewClient(oauth2.NoContext, oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_API_KEY")},
	)))

	opt := &github.ListOptions{PerPage: 100, Page: 1}
	for {
		tags, resp, err := client.Repositories.ListTags(owner, repo, opt)
		if err != nil {
			return allVersions, err
		}
		for _, tag := range tags {
			version, err := version.NewVersion(*tag.Name)
			if err == nil {
				allVersions = append(allVersions, version)
			} else {
				fmt.Fprintln(os.Stderr, fmt.Sprintf("event=SKIPPING_TAG owner=%s repo=%s tag='%s' reason='%s'", owner, repo, *tag.Name, err.Error()))
			}
		}
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	if len(allVersions) == 0 {
		return allVersions, errors.New(fmt.Sprintf("%s/%s has no valid semver releases on GitHub", owner, repo))
	}
	return allVersions, nil
}
