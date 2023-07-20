package implement_test

import (
	"fmt"
	"testing"

	gradlep "github.com/JitenPalaparthi/depscan/implement/gradle"
	pipp "github.com/JitenPalaparthi/depscan/implement/pip"
)

func TestPip(t *testing.T) {
	pip := new(pipp.Pip)
	pip.FilePaths = append(pip.FilePaths, "requirements.txt")
	gdeps, err := pip.Scan()
	fmt.Println(gdeps, err)
	if err != nil {
		t.Fail()
	}
	if len(gdeps) != 11 {
		t.Fail()
	}
}

func TestGradle(t *testing.T) {
	gradle := new(gradlep.Gradle)
	gradle.FilePaths = append(gradle.FilePaths, "dependencies.lock")
	gdeps, err := gradle.Scan()
	fmt.Println(gdeps, err)
	if err != nil {
		t.Fail()
	}
}
