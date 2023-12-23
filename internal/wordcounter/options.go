package wordcounter

import "github.com/ikawaha/kagome/v2/filter"

type Options struct {
	stopPOSList []filter.POS
}

type Option func(*Options)

func StopPOSList(posList ...POS) Option {
	return func(options *Options) {
		for _, pos := range posList {
			options.stopPOSList = append(options.stopPOSList, pos.ToFilter())
		}
	}
}
