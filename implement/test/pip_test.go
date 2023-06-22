package implement_test

import (
	"fmt"
	"testing"

	"github.com/JitenPalaparthi/depscan/implement"
)

func TestXxx(t *testing.T) {
	pip := new(implement.Pip)
	pip.FilePath = "requirements.txt"
	gdeps, err := pip.Scan()
	fmt.Println(gdeps, err)
	if err != nil {
		t.Fail()
	}
	if len(gdeps) != 11 {
		t.Fail()
	}
}
