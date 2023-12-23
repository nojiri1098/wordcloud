package wordcounter

import "github.com/ikawaha/kagome/v2/filter"

type POS []string

func (p POS) ToFilter() filter.POS {
	return filter.POS(p)
}
