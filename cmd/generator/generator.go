package main

import (
	"os"

	"github.com/nojiri1098/wordcloud/internal/wordcloud"
	"github.com/nojiri1098/wordcloud/internal/wordcounter"
)

func main() {
	// 用途に応じて特定の品詞を除外できる
	stopPOSList := wordcounter.StopPOSList([]wordcounter.POS{
		{"助詞"},
		{"助動詞"},
		{"記号"},
		{"連体詞"},
		{"副詞", "助詞類接続"},
		{"動詞", "非自立"},
		{"動詞", "接尾"},
		{"名詞", "代名詞"},
		{"名詞", "非自立"},
		{"名詞", "接尾"},
		{"名詞", "数"},
		{"名詞", "サ変接続"},
		{"フィラー"},
	}...)

	// ノイズになる単語を除外できる
	stopWords := wordcounter.StopWords(
		"ある",
		"ない",
		"いい",
		"よく",
		"どう",
		"あっ",
		"し",
		"する",
		"なる",
		"できる",
	)

	// 特定の品詞だけを抽出できる
	keepPOSList := wordcounter.KeepPOSList([]wordcounter.POS{
		{"名詞"},
		{"カスタム名詞"},
	}...)

	// カスタム名詞を追加できる
	userDict := wordcounter.UserDict("user_dict.txt")

	counter, err := wordcounter.New(
		stopPOSList,
		stopWords,
		keepPOSList,
		userDict,
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

	wordList, err := counter.Count(f)
	if err != nil {
		panic(err)
	}

	saveAs := "wordcloud"

	if err := wordcloud.New(wordList).SaveAsPNG(saveAs); err != nil {
		panic(err)
	}
}
