package task2

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func Test_Task2(t *testing.T) {
	tests := []struct {
		name  string
		testA int
		testB int
		testC int
		want  string
	}{
		{
			name:  "Zeroes",
			testA: 0,
			testB: 0,
			testC: 0,
			want:  "YES",
		},

		{
			name:  "3_4_5",
			testA: 3,
			testB: 4,
			testC: 5,
			want:  "YES",
		},

		{
			name:  "6_8_10",
			testA: 6,
			testB: 8,
			testC: 10.0,
			want:  "YES",
		},

		{
			name:  "42_1_1",
			testA: 1,
			testB: 42,
			testC: 1,
			want:  "NO",
		},

		{
			name:  "Invalid data",
			testA: -1,
			testB: 10,
			testC: 43,
			want:  "NO",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := task2(test.testA, test.testB, test.testC)
			assert.Equal(t, res, test.want)
		})
	}
}
