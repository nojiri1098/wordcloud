package main

import (
	"os"

	"github.com/nojiri1098/wordcloud/internal/wordcloud"
	"github.com/nojiri1098/wordcloud/internal/wordcounter"
)

func main() {
	counter, err := wordcounter.New(
		wordcounter.ConfigPath("../../config.yml"),
	)
	if err != nil {
		panic(err)
	}

	// 解析する対象を指定する
	// io.Reader であればなんでも良い
	f, err := os.Open("20231221.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// TODO: テキストのクレンジング

	wordList, err := counter.Count(f)
	if err != nil {
		panic(err)
	}

	saveAs := "wordcloud"

	if err := wordcloud.New(wordList).SaveAsPNG(saveAs); err != nil {
		panic(err)
	}
}
