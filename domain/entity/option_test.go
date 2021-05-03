package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListLimitOption_Apply(t *testing.T) {
	tests := []struct {
		name string
		in   ListLimitOption
		out  int
	}{
		{
			name: "success",
			in:   1,
			out:  1,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			opts := &ListOptions{}
			tt.in.Apply(opts)
			assert.Equal(t, tt.out, opts.Limit)
		})
	}
}

func Test_WithLimit(t *testing.T) {
	tests := []struct {
		name string
		in   int
		out  ListLimitOption
	}{
		{
			name: "success",
			in:   1,
			out:  ListLimitOption(1),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			out := WithLimit(tt.in)
			assert.Equal(t, tt.out, out)
		})
	}
}

func TestListOffsetOption_Apply(t *testing.T) {
	tests := []struct {
		name string
		in   ListOffsetOption
		out  int
	}{
		{
			name: "success",
			in:   1,
			out:  1,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			opts := &ListOptions{}
			tt.in.Apply(opts)
			assert.Equal(t, tt.out, opts.Offset)
		})
	}
}

func Test_WithOffset(t *testing.T) {
	tests := []struct {
		name string
		in   int
		out  ListOffsetOption
	}{
		{
			name: "success",
			in:   1,
			out:  ListOffsetOption(1),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			out := WithOffset(tt.in)
			assert.Equal(t, tt.out, out)
		})
	}
}
