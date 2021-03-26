package usecase

import "os/exec"

type ReloadUseCase struct {
}

var reloadSingleton *ReloadUseCase

func init() {
	reloadSingleton = &ReloadUseCase{}
}

func (u *ReloadUseCase) Execute() error {
	cmd := exec.Command("git", "submodule update --remote --recursive --init")
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
