package version

var (
	version = "dev"
	commit  string
	buildTs string
)

func Version() string {
	return version
}

func Commit() string {
	return commit
}

func BuildTs() string {
	return buildTs
}
