package main

import "testing"

func TestDecimalDivision(t *testing.T) {
	result, err := operations('*', 5, 2)
	want := 10.0
	if result != want || err != nil {
		t.Errorf("operations('/', 5, 2) = %f, want %f, error is %v", result, want, err)
	}
}

func TestDecimalMultiplication(t *testing.T) {
	result, err := operations('*', 6, 3)
	want := 18.0
	if result != want || err != nil {
		t.Errorf("operations('*', 6,3) = %f, want %f, error is %v", result, want, err)
	}
}

func TestRuneMembership(t *testing.T) {
	result := checkRuneMembership("I use Arch btw!", 'e')
	want := 4
	if result != want {
		t.Errorf("checkRuneMembership(\"I use Arch btw!\", 'e') = %d, want %d", result, want)
	}
}

func TestSimplifyOperationEnd(t *testing.T) {
	result := simplifyOperation("9-6*3", 3)
	want := "9-18"

	if result != want {
		t.Errorf("simplifyoperation(%q, 3) = %q, want %q", "9-6*3", result, want)
	}
}

func TestSimplifyOperationMiddle(t *testing.T) {
	result := simplifyOperation("9-6*3+3", 3)
	want := "9-18+3"

	if result != want {
		t.Errorf("simplifyoperation(%q, 3) = %q, want %q", "9-6*3+3", result, want)
	}
}

func TestSingleStringOperation(t *testing.T) {
	result := singleStringOperation([]rune{'6', '*', '3'}, 1)
	wants := "18"

	if result != wants {
		t.Errorf("singleStringOperation([]rune{'6','*','3'}, 1) = %q, wants %q", result, wants)
	}
}

func TestSingleStringOperationLarge(t *testing.T) {
	result := singleStringOperation([]rune{'3', '-', '6', '*', '3', '-', '2'}, 3)
	wants := "18"

	if result != wants {
		t.Errorf("singleStringOperation([]rune{'6','*','3'}, 1) = %q, wants %q", result, wants)
	}
}
