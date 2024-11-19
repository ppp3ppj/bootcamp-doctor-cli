package domain

type Package struct {
    Name string
    Command string
}

type PackageManager interface {
	GetAll() []Package
	Get(name string) (Package, bool)
	Create(pkg Package) error
	Update(name string, pkg Package) error
	Delete(name string) error
}
