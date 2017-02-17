package golangvectorspace

import (
	"testing"
)

func TestConcordance(t *testing.T) {
	var concordance = Concordance("this is a test")

	got := concordance["this"]
	if got != 1 {
		t.Errorf("Cocordance expect (test) to have 1")
	}

	got = concordance["is"]
	if got != 1 {
		t.Errorf("Cocordance expect (test) to have 1")
	}

	got = concordance["a"]
	if got != 1 {
		t.Errorf("Cocordance expect (test) to have 1")
	}

	got = concordance["test"]
	if got != 1 {
		t.Errorf("Cocordance expect (test) to have 1")
	}
}

func TestMagnitude(t *testing.T) {
	var concordance = Concordance("this is a test")
	var got = Magnitude(concordance)

	if got != 2 {
		t.Errorf("Magnitude expected 2")
	}

}