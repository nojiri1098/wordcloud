package main

import (
	"flag"
	"os"

	"github.com/nojiri1098/wordcloud/internal/wordcloud"
	"github.com/nojiri1098/wordcloud/internal/wordcounter"
)

var flags = struct {
	config *string
	input  *string
	output *string
}{}

func init() {
	flags.config = flag.String("config", "config.yml", "config file path")
	flags.input = flag.String("input", "cmd/generator/20231221.txt", "input file path")
	flags.output = flag.String("output", "wordcloud.png", "output file path")
}

func main() {
	flag.Parse()

	var opt func(*wordcounter.Options)
	if flags.config != nil {
		opt = wordcounter.ConfigPath(*flags.config)
	}

	counter, err := wordcounter.New(opt)
	if err != nil {
		panic(err)
	}

	f, err := os.Open(*flags.input)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// TODO: テキストのクレンジング

	wordList, err := counter.Count(f)
	if err != nil {
		panic(err)
	}

	if err := wordcloud.New(wordList).SaveAsPNG(*flags.output); err != nil {
		panic(err)
	}
}
