package application

import (
	"errors"

	domain "github.com/ppp3ppj/bootcamp-doctor-cli/domain/package"
	"github.com/spf13/viper"
)


type PackageManagerImpl struct{}

func (pm *PackageManagerImpl) GetAll() []domain.Package {
	var packages []domain.Package
	err := viper.UnmarshalKey("packages", &packages)
	if err != nil {
		return nil
	}
	return packages
}

func (pm *PackageManagerImpl) Get(name string) (domain.Package, bool) {
	packages := pm.GetAll()
	for _, pkg := range packages {
		if pkg.Name == name {
			return pkg, true
		}
	}
	return domain.Package{}, false
}

func (pm *PackageManagerImpl) Create(pkg domain.Package) error {
	packages := pm.GetAll()
	for _, p := range packages {
		if p.Name == pkg.Name {
			return errors.New("package already exists")
		}
	}
	packages = append(packages, pkg)
	viper.Set("packages", packages)
	return viper.WriteConfig()
}

func (pm *PackageManagerImpl) Update(name string, pkg domain.Package) error {
	packages := pm.GetAll()
	for i, p := range packages {
		if p.Name == name {
			packages[i] = pkg
			viper.Set("packages", packages)
			return viper.WriteConfig()
		}
	}
	return errors.New("package not found")
}

func (pm *PackageManagerImpl) Delete(name string) error {
	packages := pm.GetAll()
	for i, p := range packages {
		if p.Name == name {
			packages = append(packages[:i], packages[i+1:]...)
			viper.Set("packages", packages)
			return viper.WriteConfig()
		}
	}
	return errors.New("package not found")
}

