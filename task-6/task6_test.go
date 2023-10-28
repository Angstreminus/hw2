package task6

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestTask6(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{
			name: "0 1 2",
			arr:  []int{1, 2, 3, 2, 1},
			want: 3,
		},

		{
			name: "0 0 0",
			arr:  []int{0, 0, 0},
			want: 1,
		},

		{
			name: "1-10",
			arr:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want: 10,
		},

		{
			name: "1 2 3 4 5 1 2 1 2 7 3",
			arr:  []int{1, 2, 3, 4, 5, 1, 2, 1, 2, 7, 3},
			want: 6,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := task6(test.arr)
			assert.Equal(t, res, test.want)
		})
	}
}
