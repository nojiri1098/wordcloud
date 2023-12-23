package wordcounter

import (
	"bufio"
	"io"
	"strings"

	"github.com/ikawaha/kagome-dict/ipa"
	"github.com/ikawaha/kagome/v2/tokenizer"
)

type WordCounter struct {
	tokenizer *tokenizer.Tokenizer
}

func New() (*WordCounter, error) {
	t, err := tokenizer.New(ipa.Dict())
	if err != nil {
		return nil, err
	}

	return &WordCounter{
		tokenizer: t,
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