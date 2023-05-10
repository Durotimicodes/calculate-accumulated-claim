package testfiles

import (
	"testing"

	"github.com/Durotimicodes/wtw-intern-developer/app"
	"github.com/Durotimicodes/wtw-intern-developer/helperfunctions"
)

func TestClaimApp(t *testing.T) { // testing main program
	tests := []struct {
		testName  string
		input     string
		outputOne [][]string
	}{
		{testName: "Test-Case-One", input: "test.txt", outputOne: [][]string{
			{"1990", "4"},
			{"Comp"}, {"0.00"}, {"0.00"}, {"0.00"}, {"0.00"}, {"0.00"}, {"0.00"}, {"0.00"}, {"110.00"}, {"280.00"}, {"200.00"},
			{"Non-Comp"}, {"45.20"}, {"110.00"}, {"110.00"}, {"147.00"}, {"50.00"}, {"125.00"}, {"150.00"}, {"55.00"}, {"140.00"}, {"100.00"},
		}},
		{testName: "Test-Case-Two", input: "testTwo.txt", outputOne: [][]string{
			{"1990", "4"},
			{"Comp"}, {"0.00"}, {"0.00"}, {"0.00"}, {"0.00"}, {"0.00"}, {"0.00"}, {"0.00"}, {"110.00"}, {"280.00"}, {"200.00"},
			{"Non-Comp"}, {"45.20"}, {"110.00"}, {"110.00"}, {"147.00"}, {"0.00"}, {"0.00"}, {"0.00"}, {"0.00"}, {"85.00"}, {"100.00"},
			{"Mech-Product"}, {"0.00"}, {"0.00"}, {"0.00"}, {"0.00"}, {"100.00"}, {"200.00"}, {"310.00"}, {"0.00"}, {"0.00"}, {"0.00"},
		}},
		{testName: "Test-Case-Three", input: "testThree.txt", outputOne: [][]string{
			{"1990", "4"},
			{"Comp"}, {"0.00"}, {"0.00"}, {"0.00"}, {"0.00"}, {"0.00"}, {"0.00"}, {"0.00"}, {"110.00"}, {"280.00"}, {"200.00"},
			{"Non-Comp"}, {"45.20"}, {"110.00"}, {"110.00"}, {"147.00"}, {"0.00"}, {"0.00"}, {"0.00"}, {"0.00"}, {"85.00"}, {"100.00"},
			{"Mech-Product"}, {"0.00"}, {"0.00"}, {"0.00"}, {"0.00"}, {"100.00"}, {"200.00"}, {"310.00"}, {"0.00"}, {"0.00"}, {"0.00"},
		}},
	}

	for _, tc := range tests {
		gotOne := app.ClaimApp(tc.input)

		if helperfunctions.CheckDataAsStrings(gotOne, tc.outputOne) {
			t.Fatalf("name : %s, expectedOne: %v, gotOne : %v", tc.testName, tc.outputOne, gotOne)
		}

	}
}
