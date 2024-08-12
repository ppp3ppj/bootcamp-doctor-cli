package resources

import (
	"fmt"
	"math/rand"
)

var packages = []string{
    "minji",
    "golang",
}

func GetPackages() []string {
    pkgs := packages
    copy(pkgs, packages)

	rand.Shuffle(len(pkgs), func(i, j int) {
		pkgs[i], pkgs[j] = pkgs[j], pkgs[i]
	})

    for k := range pkgs {
		pkgs[k] += fmt.Sprintf("-%d.%d.%d", rand.Intn(10), rand.Intn(10), rand.Intn(10)) //nolint:gosec
    }

    return pkgs
}
