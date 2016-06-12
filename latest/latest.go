package latest

import (
	"github.com/matthew-andrews/go-version"
	"sort"
)

func Latest(versions []*version.Version) string {
	sort.Sort(version.Collection(versions))
	return versions[len(versions)-1].Raw()
}
