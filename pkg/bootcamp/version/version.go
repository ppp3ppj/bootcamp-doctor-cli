package version

import "fmt"

var (
    Version = "dev"
    Hash = "0.0.1"
)

func Print() string {
    return fmt.Sprintf("%s-%s", Hash, Version)
}
