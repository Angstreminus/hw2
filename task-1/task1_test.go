package task1

import (
	"errors"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestValidateInput(t *testing.T) {
	tests := []struct {
		name  string
		testA int
		testB int
		want  error
	}{
		{
			name:  "Border_InvalidTest_A",
			testA: -1,
			testB: 10,
			want:  errors.New("INVALID INPUT"),
		},
		{
			name:  "Border_InvalidTest_A",
			testA: 1001,
			testB: 10,
			want:  errors.New("INVALID INPUT"),
		},
		{
			name:  "Border_InvalidTest_B",
			testA: 10,
			testB: -1,
			want:  errors.New("INVALID INPUT"),
		},
		{
			name:  "Border_InvalidTest_B",
			testA: 10,
			testB: 1001,
			want:  errors.New("INVALID INPUT"),
		},
		{
			name:  "Border_InvalidTest_BothArgs",
			testA: -1,
			testB: -1,
			want:  errors.New("INVALID INPUT"),
		},
		{
			name:  "ValidTest_Input",
			testA: 3,
			testB: 2,
			want:  nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := validateInput(test.testA, test.testB)
			assert.Equal(t, test.want, err)
		})
	}
}

func TestTask1(t *testing.T) {
	tests := []struct {
		name    string
		testA   int
		testB   int
		want    float64
		wantErr error
	}{
		{
			name:    "Zeroes",
			testA:   0,
			testB:   0,
			want:    0.0,
			wantErr: nil,
		},
		{
			name:    "3_4_5",
			testA:   3,
			testB:   4,
			want:    5.0,
			wantErr: nil,
		},
		{
			name:    "6_8_10",
			testA:   6,
			testB:   8,
			want:    10.0,
			wantErr: nil,
		},
		{
			name:    "Invalid data",
			testA:   -1,
			testB:   10,
			want:    0.0,
			wantErr: errors.New("INVALID INPUT"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res, err := task1(test.testA, test.testB)
			assert.Equal(t, test.want, res)
			assert.Equal(t, test.wantErr, err)
		})
	}
}
