# Semver as a Service

Similar to [semver.io](https://semver.io) but hasn't got all the same features (yet) but supports all repositories on GitHub rather than a small curated selection of projects.

Current feature set:-

- Get the latest (the highest semver tag) of tags pushed to GitHub for any repository
- The tag is returned in the same format as the original tag so that install scripts can be easily written, e.g.

## Examples

Tell me the latest version of [Hugo](https://github.com/spf13/hugo/):-

```
curl -f https://api.mattandre.ws/semver/github/spf13/hugo
```

The idea is that this API can be combined with other scripts, for example install scripts.

The following snippet installs the latest version of [apex](https://github.com/apex/apex) for OS X:-

```
curl -sf https://api.mattandre.ws/semver/github/apex/apex \
	| xargs -I '{}' echo https://github.com/apex/apex/releases/{}/download/apex_darwin_386 -o /usr/local/bin/apex \
	&& chmod +x /usr/local/bin/apex
```
