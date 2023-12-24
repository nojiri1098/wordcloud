package wordcounter

import (
	"bufio"
	"io"
	"strings"

	"github.com/ikawaha/kagome-dict/dict"
	"github.com/ikawaha/kagome-dict/ipa"
	"github.com/ikawaha/kagome/v2/filter"
	"github.com/ikawaha/kagome/v2/tokenizer"
)

type WordCounter struct {
	tokenizer *tokenizer.Tokenizer
	options   *Options
}

func New(options ...Option) (*WordCounter, error) {
	counterOptions := &Options{}
	for _, option := range options {
		option(counterOptions)
	}

	if path := counterOptions.ConfigPath; path != "" {
		config, err := LoadConfig(path)
		if err != nil {
			return nil, err
		}

		for _, pos := range config.ExcludePOSList {
			counterOptions.excludePOSList = append(counterOptions.excludePOSList, pos.ToFilter())
		}

		for _, pos := range config.KeepPOSList {
			counterOptions.keepPOSList = append(counterOptions.keepPOSList, pos.ToFilter())
		}

		counterOptions.stopWords = config.StopWords

		counterOptions.threshold = config.Threshold

		counterOptions.userDict = config.UserDict
	}

	tokenizerOptions := []tokenizer.Option{
		tokenizer.OmitBosEos(),
	}
	if len(counterOptions.userDict) > 0 {
		s := strings.NewReader(strings.Join(counterOptions.userDict, "\n"))
		r, err := dict.NewUserDicRecords(s)
		if err != nil {
			return nil, err
		}
		userDict, err := r.NewUserDict()
		if err != nil {
			return nil, err
		}

		tokenizerOptions = append(tokenizerOptions, tokenizer.UserDict(userDict))
	}

	t, err := tokenizer.New(ipa.Dict(), tokenizerOptions...)
	if err != nil {
		return nil, err
	}

	return &WordCounter{
		tokenizer: t,
		options:   counterOptions,
	}, nil
}

func (wc *WordCounter) Count(r io.Reader) (map[string]int, error) {
	// read lines
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// tokenize
	text := strings.Join(lines, " ")
	tokens := wc.tokenizer.Tokenize(text)

	// filter tokens
	filter.NewPOSFilter(wc.options.excludePOSList...).Drop(&tokens)
	filter.NewWordFilter(wc.options.stopWords).Drop(&tokens)
	filter.NewPOSFilter(wc.options.keepPOSList...).Keep(&tokens)

	// count words
	result := make(map[string]int)
	for _, token := range tokens {
		if token.Class == tokenizer.DUMMY {
			continue
		}

		b, ok := token.BaseForm()
		if !ok {
			result[token.Surface]++
			continue
		}

		result[b]++
	}

	// apply threshould
	for word, count := range result {
		if count < wc.options.threshold {
			delete(result, word)
		}
	}

	return result, nil
}
