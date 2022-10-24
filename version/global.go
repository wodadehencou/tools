package version

import "fmt"

var (
	Name  string
	Major int
	Minor int
	Patch int
	Hash  string
)

func Full() string {
	s := fmt.Sprintf("%s-v%d.%d.%d", Name, Major, Minor, Patch)
	if len(Hash) != 0 {
		s = fmt.Sprintf("%s-%s", s, Hash)
	}
	return s
}

func Short() string {
	s := fmt.Sprintf("%d.%d.%d", Major, Minor, Patch)
	if len(Hash) != 0 {
		s = fmt.Sprintf("%s-%s", s, Hash)
	}
	return s
}
