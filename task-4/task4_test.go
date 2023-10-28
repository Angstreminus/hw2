package task4

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func Test_Task4(t *testing.T) {
	tests := []struct {
		name  string
		input [][]int
		want  bool
	}{
		{
			name:  "true",
			input: [][]int{{0, 1, 2}, {1, 5, 3}, {2, 3, 4}},
			want:  true,
		},

		{
			name:  "false",
			input: [][]int{{0, 0, 0}, {0, 0, 0}, {1, 0, 0}},
			want:  false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := task4(test.input)
			assert.Equal(t, res, test.want)
		})
	}
}
