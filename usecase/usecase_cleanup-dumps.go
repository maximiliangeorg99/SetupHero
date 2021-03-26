package usecase

import "os/exec"

type CleanUpDumpsUseCase struct {
}

var cleanUpDumpsSingleton *CleanUpDumpsUseCase

func init() {
	cleanUpDumpsSingleton = &CleanUpDumpsUseCase{}
}

func (u *CleanUpDumpsUseCase) Execute() error {
	cmd := exec.Command("git", "submodule foreach -recursive rm -rf dumps")
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
