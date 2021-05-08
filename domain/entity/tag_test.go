package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTagList_GetIDs(t *testing.T) {
	tests := []struct {
		name string
		in   TagList
		out  []int
	}{
		{
			name: "success",
			in: TagList{
				{ID: 1},
				{ID: 2},
			},
			out: []int{1, 2},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			out := tt.in.GetIDs()
			assert.Equal(t, tt.out, out)
		})
	}
}
