package application

import domain "github.com/ppp3ppj/bootcamp-doctor-cli/domain/cli-version"

type CLIVersionUseCase struct {
	Version domain.CLIVersion
}

func (v *CLIVersionUseCase) GetCurrentVersion() domain.CLIVersion {
	return v.Version
}
