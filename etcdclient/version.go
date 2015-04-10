package etcdclient

type Version struct {
	ReleaseVersion  string
	InternalVersion string
}

func (v *Version) String() string {
	return v.ReleaseVersion
}
