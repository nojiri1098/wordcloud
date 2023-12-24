package wordcloud

import (
	"image/color"
	"image/png"
	"os"
	"strings"

	"github.com/psykhi/wordclouds"
)

type WordCloud struct {
	client *wordclouds.Wordcloud
}

var presetColors = []color.RGBA{
	{0x00, 0x00, 0x00, 0xff},
	{0x46, 0x46, 0x46, 0xff},
	{0x64, 0x64, 0x64, 0xff},
	{0xbb, 0xbb, 0xbb, 0xff},
	{0xc8, 0xc8, 0xc8, 0xff},
}

func New(wordList map[string]int) WordCloud {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	println(wd)

	return WordCloud{
		client: wordclouds.NewWordcloud(
			wordList,
			wordclouds.Colors(func() (colors []color.Color) {
				for _, c := range presetColors {
					colors = append(colors, c)
				}
				return
			}()),
			wordclouds.FontFile("internal/wordcloud/font/BIZUDPGothic-Regular.ttf"),
		),
	}
}

func (wc WordCloud) SaveAsPNG(path string) error {
	if !strings.HasSuffix(path, ".png") {
		path = path + ".png"
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := png.Encode(f, wc.client.Draw()); err != nil {
		return err
	}

	return nil
}
