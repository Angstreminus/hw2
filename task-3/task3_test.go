package task3

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func Test_Task3(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "1 2 3",
			input: []int{1, 2, 3},
			want:  []int{3, 1, 2},
		},
		{
			name:  "4, 5, 3, 4, 2, 3",
			input: []int{4, 5, 3, 4, 2, 3},
			want:  []int{3, 4, 5, 3, 4, 2},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := task3(test.input)
			assert.Equal(t, test.want, res)
		})
	}
}
