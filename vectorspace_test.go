package golangvectorspace

import (
	"testing"
)

func TestConcordance(t *testing.T) {
	var concordance = Concordance("this is a test")

	got := concordance["this"]
	if got != 1 {
		t.Errorf("Cocordance expect (test) to have 1 but got %q", got)
	}

	got = concordance["is"]
	if got != 1 {
		t.Errorf("Cocordance expect (test) to have 1 but got %q", got)
	}

	got = concordance["a"]
	if got != 1 {
		t.Errorf("Cocordance expect (test) to have 1 but got %q", got)
	}

	got = concordance["test"]
	if got != 1 {
		t.Errorf("Cocordance expect (test) to have 1 but got %q", got)
	}
}

func TestMagnitude(t *testing.T) {
	var concordance = Concordance("this is a test")
	var got = Magnitude(concordance)

	if got != 2 {
		t.Errorf("Magnitude expected 2 but got %q", got)
	}

}

func TestRelationSameExpectOne(t *testing.T) {
	var concordance1 = Concordance("this is a test")
	var concordance2 = Concordance("this is a test")

	got := Relation(concordance1, concordance2)

	if got != 1 {
		t.Errorf("Relation expected 1 but got %q", got)
	}
}

func TestRelationSimilarExpectSimilar(t *testing.T) {
	var concordance1 = Concordance("this is a test")
	var concordance2 = Concordance("this test")

	got := Relation(concordance1, concordance2)

	if got != 0.7071067811865475 {
		t.Errorf("Relation expected 0.7071067811865475 but got %q", got)
	}
}

func TestRelationDifferentExpectZero(t *testing.T) {
	var concordance1 = Concordance("this is a test")
	var concordance2 = Concordance("not related at all")

	got := Relation(concordance1, concordance2)

	if got != 0 {
		t.Errorf("Relation expected 0.5 but got %q", got)
	}
}

func TestRelationSimilarStrings(t *testing.T) {
	var concordance1 = Concordance("Go has a lightweight test framework composed of the go test command and the testing package.")
	var concordance2 = Concordance("Package testing provides support for automated testing of Go packages. It is intended to be used in concert with the go test command, which automates execution of any function of the form.")

	got := Relation(concordance1, concordance2)

	if got != 0.48211825989991874 {
		t.Errorf("Relation expected 0.48211825989991874 but got %q", got)
	}
}