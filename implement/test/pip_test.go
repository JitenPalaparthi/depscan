package implement_test

import (
	"fmt"
	"testing"

	"github.com/JitenPalaparthi/depscan/implement"
)

func TestXxx(t *testing.T) {
	pip := new(implement.Pip)
	gdeps, err := pip.Scan("requirements.txt")
	fmt.Println(gdeps, err)
	if err != nil {
		t.Fail()
	}
	if len(gdeps) != 11 {
		t.Fail()
	}
}
