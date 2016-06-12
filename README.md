# Semver as a service

Similar to [semver.io](https://semver.io) but hasn't got all the same features (yet) and support all repositories on GitHub rather than a small curated selection of projects.

Currently only supports getting the latest (the highest semver tags) of tags pushed to GitHub.

## Examples

Latest [Hugo](https://github.com/spf13/hugo/):-

```
curl -f https://api.mattandre.ws/semver/github/spf13/hugo
```

Latest [denodeify](https://github.com/matthew-andrews/denodeify)

```
curl -f https://api.mattandre.ws/semver/matthew-andrews/denodeify
```
