package main

import "github.com/nojiri1098/wordcloud/internal/wordcloud"

func main() {
	wordList := map[string]int{
		"word": 1,
	}

	saveAs := "wordcloud"

	if err := wordcloud.New(wordList).SaveAsPNG(saveAs); err != nil {
		panic(err)
	}
}
