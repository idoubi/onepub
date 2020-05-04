package util

import (
	"errors"
	"github.com/russross/blackfriday"
	"io/ioutil"
	"os"
	"strings"
)

// 博客信息
type Article struct {
	Title       string
	Content     string
	HtmlContent string
}

func NewArticleByMdFile(mdfilePath string) (article Article, err error) {
	fileInfo, err := os.Stat(mdfilePath)
	// 从文件名中获取标题
	if err != nil {
		return
	}

	info := strings.Split(fileInfo.Name(), ".")
	if info[1] != "md" {
		err = errors.New("NewArticleByMdFile: " + mdfilePath + " not a md file")
		return
	}

	mdTitle := info[0]

	mdContent, err := ioutil.ReadFile(mdfilePath)
	if err != nil {
		return
	}

	output := blackfriday.Run(mdContent)

	return Article{
		Title:       mdTitle,
		Content:     string(mdContent),
		HtmlContent: string(output),
	}, nil
}
