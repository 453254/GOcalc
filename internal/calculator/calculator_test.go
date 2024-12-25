package calcul_test

import (
	"testing"

	"github.com/MaxMindshtorm/calculator/pkg/calcul"
)

func TestCalc(t *testing.T) {
	tests := []struct {
		expression string
		expected   float64
		expectErr  bool
	}{
		{"5 + 7", 12, false},
		{"12 - 4 * 2", 4, false},
		{"(1 + 4) * 3", 15, false},
		{"8 / 2", 4, false},
		{"8 / 0", 0, true},
		{"4 +", 0, true},
		{"x + 5", 0, true},
	}

	for _, test := range tests {
		result, err := calcul.Calc(test.expression)
		if (err != nil) != test.expectErr {
			t.Errorf("Calc(%q): unexpected error status: got %v, want error: %v", test.expression, err != nil, test.expectErr)
		}
		if !test.expectErr && result != test.expected {
			t.Errorf("Calc(%q): got %v, want %v", test.expression, result, test.expected)
		}
	}
}