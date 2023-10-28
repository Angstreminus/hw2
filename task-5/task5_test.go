package task5

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func Test_Task5(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "1+1",
			input: "1+1=2",
			want:  "one+one=2",
		},

		{
			name:  "2+2",
			input: "2+2=4",
			want:  "2+2=4",
		},

		{
			name:  "11+1",
			input: "11+1=12",
			want:  "oneone+one=one2",
		},

		{
			name:  "Empty",
			input: "",
			want:  "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := task5(test.input)
			assert.Equal(t, res, test.want)
		})
	}
}
