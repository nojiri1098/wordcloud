package main

import (
	"strings"

	"github.com/nojiri1098/wordcloud/internal/wordcloud"
	"github.com/nojiri1098/wordcloud/internal/wordcounter"
)

func main() {
	counter, err := wordcounter.New()
	if err != nil {
		panic(err)
	}

	r := strings.NewReader("word")
	wordList, err := counter.Count(r)
	if err != nil {
		panic(err)
	}

	saveAs := "wordcloud"

	if err := wordcloud.New(wordList).SaveAsPNG(saveAs); err != nil {
		panic(err)
	}
}
