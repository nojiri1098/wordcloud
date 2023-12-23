package wordcounter

import "github.com/ikawaha/kagome/v2/filter"

type Options struct {
	stopPOSList []filter.POS
	stopWords   []string
}

type Option func(*Options)

func StopPOSList(posList ...POS) Option {
	return func(options *Options) {
		for _, pos := range posList {
			options.stopPOSList = append(options.stopPOSList, pos.ToFilter())
		}
	}
}

func StopWords(words ...string) Option {
	return func(options *Options) {
		options.stopWords = append(options.stopWords, words...)
	}
}
