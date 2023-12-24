package wordcounter

import "github.com/ikawaha/kagome/v2/filter"

type Options struct {
	excludePOSList []filter.POS
	stopWords      []string
	keepPOSList    []filter.POS
	userDict       []string
	threshold      int
}

type Option func(*Options)

func ExcludePOSList(posList ...POS) Option {
	return func(options *Options) {
		for _, pos := range posList {
			options.excludePOSList = append(options.excludePOSList, pos.ToFilter())
		}
	}
}

func StopWords(words ...string) Option {
	return func(options *Options) {
		options.stopWords = append(options.stopWords, words...)
	}
}

func KeepPOSList(posList ...POS) Option {
	return func(options *Options) {
		for _, pos := range posList {
			options.keepPOSList = append(options.keepPOSList, pos.ToFilter())
		}
	}
}

func UserDict(userDict []string) Option {
	return func(options *Options) {
		options.userDict = userDict
	}
}

func Threshold(threshold int) Option {
	return func(options *Options) {
		options.threshold = threshold
	}
}
