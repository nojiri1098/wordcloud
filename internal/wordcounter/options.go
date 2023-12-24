package wordcounter

import "github.com/ikawaha/kagome/v2/filter"

type Options struct {
	ConfigPath string

	excludePOSList []filter.POS
	stopWords      []string
	keepPOSList    []filter.POS
	userDict       []string
	threshold      int
}

type Option func(*Options)

func ConfigPath(path string) Option {
	return func(options *Options) {
		options.ConfigPath = path
	}
}
