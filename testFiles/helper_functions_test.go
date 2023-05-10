package testfiles

import (
	"reflect"
	"testing"

	"github.com/Durotimicodes/wtw-intern-developer/helperfunctions"
)

func TestGetMinYearAndDevelopementYears(t *testing.T) { //testing helper function

	tests := []struct {
		testName  string
		input     [][]string
		outputOne int
		outputTwo int
	}{
		{testName: "Test-Case-One", input: [][]string{{
			"comp", "1992", "1992", "110.0",
		}, {
			"comp", "1992", "1993", "170.0",
		}, {
			"comp", "1993", "1993", "200.0",
		}, {
			"non-comp", "1990", "1990", "45.2",
		}, {
			"non-comp", "1990", "1991", "64.8",
		}, {
			"non-comp", "1990", "1993", "37.0",
		}, {
			"non-comp", "1991", "1991", "50.0",
		}, {
			"non-comp", "1991", "1992", "75.0",
		}, {
			"non-comp", "1991", "1993", "25.0",
		}, {
			"non-comp", "1992", "1992", "55.0",
		}, {
			"non-comp", "1992", "1993", "85.0",
		}, {
			"non-comp", "1993", "1993", "100.0",
		}}, outputOne: 1990, outputTwo: 4},

		{testName: "Test-Case-Two", input: [][]string{{
			"comp", "1992", "1992", "110.0",
		}, {
			"comp", "1992", "1993", "170.0",
		}, {
			"comp", "1993", "1993", "200.0",
		}, {
			"non-comp", "1990", "1990", "45.2",
		}, {
			"non-comp", "1990", "1991", "64.8",
		}, {
			"non-comp", "1990", "1993", "37.0",
		}, {
			"non-comp", "1990", "1994", "50",
		}, {
			"non-comp", "1991", "1991", "50.0",
		}, {
			"non-comp", "1991", "1992", "75.0",
		}, {
			"non-comp", "1991", "1993", "25.0",
		}, {
			"non-comp", "1991", "1994", "22.0",
		}, {
			"non-comp", "1992", "1992", "55.0",
		}, {
			"non-comp", "1992", "1993", "85.0",
		}, {
			"non-comp", "1993", "1993", "100.0",
		}, {
			"non-comp", "1993", "1994", "110.0",
		}, {
			"non-comp", "1994", "1994", "140.0",
		}}, outputOne: 1990, outputTwo: 5},

		{testName: "Test-Case-Three", input: [][]string{{
			"comp", "1992", "1992", "110.0",
		}, {
			"comp", "1992", "1993", "170.0",
		}, {
			"comp", "1993", "1993", "200.0",
		}, {
			"non-comp", "1992", "1992", "55.0",
		}, {
			"non-comp", "1992", "1993", "85.0",
		}, {
			"non-comp", "1993", "1993", "100.0",
		}}, outputOne: 1992, outputTwo: 3},
	}

	for _, tc := range tests {
		gotOne, gotTwo := helperfunctions.GetMinYearAndDevelopementYears(tc.input)
		if !reflect.DeepEqual(tc.outputOne, gotOne) {
			if !reflect.DeepEqual(tc.outputOne, gotTwo) {
				t.Fatalf("name : %s, expected: %v, gotOne : %v, gotTwo : %v", tc.testName, tc.outputOne, gotOne, gotTwo)
			}
		}
	}
}
