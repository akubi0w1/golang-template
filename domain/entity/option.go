package entity

type ListOptions struct {
	Limit  int
	Offset int
}

type ListOption interface {
	Apply(o *ListOptions)
}

type ListLimitOption int

func (v ListLimitOption) Apply(o *ListOptions) {
	o.Limit = int(v)
}

func WithLimit(limit int) ListLimitOption {
	return ListLimitOption(limit)
}

type ListOffsetOption int

func (v ListOffsetOption) Apply(o *ListOptions) {
	o.Offset = int(v)
}

func WithOffset(offset int) ListOffsetOption {
	return ListOffsetOption(offset)
}
