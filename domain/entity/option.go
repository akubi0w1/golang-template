package entity

type ListOptions struct {
	Limit  int
	Offset int
}

type ListOption interface {
	Apply(o *ListOptions)
}

type ListLimitOption int

// TODO-akubi: add test
func (v ListLimitOption) Apply(o *ListOptions) {
	o.Limit = int(v)
}

// TODO-akubi: add test
func WithLimit(limit int) ListLimitOption {
	return ListLimitOption(limit)
}

type ListOffsetOption int

// TODO-akubi: add test
func (v ListOffsetOption) Apply(o *ListOptions) {
	o.Offset = int(v)
}

// TODO-akubi: add test
func WithOffset(offset int) ListOffsetOption {
	return ListOffsetOption(offset)
}
