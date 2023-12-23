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

	tokenizerOptions := []tokenizer.Option{
		tokenizer.OmitBosEos(),
	}
	if counterOptions.userDictPath != "" {
		userDict, err := dict.NewUserDict(counterOptions.userDictPath)
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

	return result, nil
}
