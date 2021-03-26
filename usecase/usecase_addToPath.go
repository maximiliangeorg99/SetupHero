package usecase

import (
	"errors"
	"os"
	"path/filepath"
)

type AddToPathUseCase struct {
}

type AddToPathRequest struct {
	firstname string
	email     string
}

var addToPathSingleton *AddToPathUseCase

func init() {
	addToPathSingleton = &AddToPathUseCase{}
}

func (u *AddToPathUseCase) Execute(request *AddToPathRequest) error {
	path, _ := os.Getwd()
	err := validateAddToPathRequest(request)
	if err != nil {
		return err
	}
	p1 := filepath.FromSlash("\\..\\tools\\sdk\\OpenJDK\\11.0.2\\windows")
	p2 := filepath.FromSlash("\\..\\tools\\maven\\3.2.5")
	p3 := filepath.FromSlash("\\..\\tools\\maven\\3.2.5\\bin\\mvn.bat")
	os.Setenv("allMailReceiver", request.firstname)
	os.Setenv("nonLiveDefaultDomainSuffix", request.email+"chrono24.net")
	os.Setenv("JAVA_HOME", path+p1)
	os.Setenv("M2_HOME", path+p2)
	os.Setenv("M2_EXECUTABLE", path+p3)
	return nil
}

func validateAddToPathRequest(request *AddToPathRequest) error {
	if request.firstname == "" {
		return errors.New("Please provide a valid firstname")
	}

	if request.email == "" {
		return errors.New("Please provide a valid email")
	}

	return nil
}
