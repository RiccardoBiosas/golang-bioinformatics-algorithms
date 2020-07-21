package algorithms

import (
	"reflect"
	"testing"
)

func TestAlgo1(t *testing.T) {
	expected := map[string]uint64{"A": 20, "C": 12, "G": 17, "T": 21}
	output := algo1("AGCTTTTCATTCTGACTGCAACGGGCAATATGTCTCTGTGTGGATTAAAAAAAGAGTGTCTGATAGCAGC")
	if eq := reflect.DeepEqual(expected, output); !eq {
		t.Errorf("Output was incorrect")
	}
}

func TestAlgo2(t *testing.T) {
	expected := "GAUGGAACUUGACUACGUAAAUU"
	output := algo2("GAUGGAACUUGACUACGUAAAUU")
	if output != expected {
		t.Errorf("Output was incorrect, got: %s, expected: %s", output, expected)
	}
}
