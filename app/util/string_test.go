package util

import (
	"testing"
)
func TestToTitle(t *testing.T) {
	one := ToTitle("SimpleCamelCase")
	if one != "Simple Camel Case" {
		t.Errorf("Whoops: " + one)
	}

	two := ToTitle("CSVFilesAreCoolButTXTRules")
	if two != "CSV Files Are Cool But TXT Rules" {
		t.Errorf("Whoops: " + two)
	}

	three := ToTitle("MediaTypes")
	if three != "Media Types" {
		t.Errorf("Whoops: " + three)
	}
}
