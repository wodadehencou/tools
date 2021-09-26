package version

import "fmt"

type Version struct {
	Name  string
	Major int
	Minor int
	Patch int
}

func New(name string, major, minor, patch int) *Version {
	return &Version{
		Name:  name,
		Major: major,
		Minor: minor,
		Patch: patch,
	}
}

func (v *Version) Full() string {
	return fmt.Sprintf("%s-v%d.%d.%d", v.Name, v.Major, v.Minor, v.Patch)
}

func (v *Version) Short() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}
