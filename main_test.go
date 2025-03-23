package main_test

import (
	"fmt"
	"testing"
)

func calculateDegrees(val *int, to *string) error {
	switch *to {
	case "F":
		*val = (*val * 9 / 5) + 32
		return nil
	case "C":
		*val = (*val - 32) * 5 / 9
		return nil
	default:
		return fmt.Errorf("invalid target degree")
	}
}

func changeNumber(n *int) {
	*n = 100
}

func Test_changeNumber(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "change number to 100",
			input:    5,
			expected: 100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := tt.input
			changeNumber(&n)
			if n != tt.expected {
				t.Errorf("changeNumber() = %v, want %v", n, tt.expected)
			}
		})
	}
}

func Test_calculateDegrees(t *testing.T) {
	tests := []struct {
		name        string
		inputVal    int
		inputTarget string
		expected    int
		wantErr     bool
	}{
		{
			name:        "celsius to fahrenheit",
			inputVal:    0,
			inputTarget: "F",
			expected:    32,
			wantErr:     false,
		},
		{
			name:        "fahrenheit to celsius",
			inputVal:    32,
			inputTarget: "C",
			expected:    0,
			wantErr:     false,
		},
		{
			name:        "invalid target degree",
			inputVal:    0,
			inputTarget: "K",
			expected:    0,
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			val := tt.inputVal
			target := tt.inputTarget
			err := calculateDegrees(&val, &target)

			if (err != nil) != tt.wantErr {
				t.Errorf("calculateDegrees() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && val != tt.expected {
				t.Errorf("calculateDegrees() = %v, want %v", val, tt.expected)
			}
		})
	}
}
