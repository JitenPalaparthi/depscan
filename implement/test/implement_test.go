package implement_test

import (
	"fmt"
	"testing"

	"github.com/JitenPalaparthi/depscan/implement"
)

func TestPip(t *testing.T) {
	pip := new(implement.Pip)
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
	gradle := new(implement.Gradle)
	gradle.FilePaths = append(gradle.FilePaths, "build.gradle")
	gdeps, err := gradle.Scan()
	fmt.Println(gdeps, err)
	if err != nil {
		t.Fail()
	}
}
