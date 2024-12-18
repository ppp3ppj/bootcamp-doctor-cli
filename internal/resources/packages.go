package resources

type PackageDoctor struct {
	Name           string
	VersionCommand string
}

var packages = []PackageDoctor{
	{"golang", "go version"},
	{"node", "node -v"},
	{"zig", "zig version"},
	{"ruby", "ruby -v"},
	{"nvm", "nvm -v"},
	{"rbenv", "rbenv -v"},
}

func GetPackages() []PackageDoctor {
	return packages
}
