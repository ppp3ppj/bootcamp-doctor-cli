package resources

type PackageDoctor struct {
    Name string
    VersionCommand string
}

var packages = []PackageDoctor{
    PackageDoctor{
        Name: "golang",
        VersionCommand: "go version",
    },
    PackageDoctor{
        Name: "node",
        VersionCommand: "node -v",
    },
    PackageDoctor{
        Name: "zig",
        VersionCommand: "zig version",
    },
    PackageDoctor{
        Name: "ruby",
        VersionCommand: "ruby -v",
    },
}

func GetPackages() []PackageDoctor {
    return packages
}
