package domain

type CLIVersion struct {
    Version string
    Build string
}

func NewCLIVersion(version, build string) CLIVersion {
    return CLIVersion{
        Version: version,
        Build: build,
    }
}
