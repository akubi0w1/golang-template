package mysql

import "github.com/akubi0w1/golang-sample/domain/entity"

func mergeListOptions(opts []entity.ListOption) *entity.ListOptions {
	opt := new(entity.ListOptions)
	for i := range opts {
		opts[i].Apply(opt)
	}
	return opt
}
