package application

import (
	"errors"
	"os/exec"
)

type PackageManager struct {
    Packages map[string]string // Key: Package name, Value: Command
}

// Add a package
func (pm *PackageManager) AddPackage(name, cmd string) error {
    if _, exists := pm.Packages[name]; exists {
        return errors.New("package already exits");
    }

    pm.Packages[name] = cmd
    return nil
}

// Get all packages
func (pm *PackageManager) ListPackages() map[string]string {
    return pm.Packages
}

// Update a package
func (pm *PackageManager) UpdatePackage(name, newCmd string) error {
    if _, exists := pm.Packages[name]; !exists {
        return errors.New("package not found")
    }

    pm.Packages[name] = newCmd
    return nil
}

// Delete a package
func (pm *PackageManager) DeletePackage(name string) error {
    if _, exists := pm.Packages[name]; !exists {
        return errors.New("package not found")
    }

    delete(pm.Packages, name)
    return nil
}

// Check the version of a package
func (pm *PackageManager) CheckVersion(name string) (string, error) {
    cmdStr, exists := pm.Packages[name]
    if !exists {
        return "", errors.New("package not found")
    }

    output, err := exec.Command("sh", "-c", cmdStr).Output()
    if err != nil {
        return "", err
    }
    return string(output), nil
}

