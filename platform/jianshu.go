package platform

import (
	"encoding/json"
	"fmt"
	"github.com/idoubi/goz"
	"github.com/idoubi/onepub/util"
	"github.com/spf13/viper"
	"strconv"
)

var (
	jianshuBlogTypeUrl    string   = "https://www.jianshu.com/author/notebooks"
	jianshuAddBlogUrl     string   = "https://www.jianshu.com/author/notes"
	jianshuEditBlogUrl    string   = "https://www.jianshu.com/author/notes/%s"
	jianshuReleaseBlogUrl string   = "https://www.jianshu.com/author/notes/%s/publicize"
	jisnshuBlogType       []string = []string{"note", "diary"}
)

// 简书 https://www.jianshu.com/
type JianShu struct {
	info platInfo
}

// Login 模拟登陆
func (blog *JianShu) Login() error {
	return nil
}

func (blog *JianShu) IsLogin() error {
	_, _, err := blog.getDiaryNoteId()
	return err
}

// 发布
func (blog *JianShu) Publish(article util.Article) error {
	noteId, err := blog.getNoteIdByConf()
	if err != nil {
		return fmt.Errorf("jisnshu get noteid error: %w", err)
	}

	articleId, err := blog.addArticle(noteId, article)
	if err != nil {
		return fmt.Errorf("jianshu add article error: %w", err)
	}

	err = blog.editArticle(articleId, article)
	if err != nil {
		return fmt.Errorf("jianshu edit article error: %w", err)
	}

	err = blog.releaseArticle(articleId)
	if err != nil {
		return fmt.Errorf("jianshu release article error: %w", err)
	}

	return nil
}

// 根据配置获取日志或随笔id
func (blog *JianShu) getNoteIdByConf() (string, error) {
	diaryId, noteId, err := blog.getDiaryNoteId()
	if err != nil {
		return "", err
	}

	var confBlogType string = viper.GetString("platform.jianshu.type")
	if !util.InSlice(confBlogType, jisnshuBlogType) {
		return "", fmt.Errorf("blog type error: %s", confBlogType)
	}

	if confBlogType == "note" {
		return noteId, nil
	} else {
		return diaryId, nil
	}
}

func (blog *JianShu) getDiaryNoteId() (diaryId, noteId string, err error) {
	resp, err := cli.Get(jianshuBlogTypeUrl, goz.Options{
		Headers: blog.headers(),
	})

	if err != nil {
		return
	}

	body, err := resp.GetBody()
	if err != nil {
		return
	}

	var respMap []map[string]interface{}
	err = json.Unmarshal(body, &respMap)
	if err != nil {
		return
	}

	for _, info := range respMap {
		if info["name"].(string) == "日记本" {
			diaryId = strconv.FormatFloat(info["id"].(float64), 'f', -1, 64)
		}

		if info["name"].(string) == "随笔" {
			noteId = strconv.FormatFloat(info["id"].(float64), 'f', -1, 64)
		}
	}

	return
}

// 新增
func (blog *JianShu) addArticle(noteId string, article util.Article) (articleId string, err error) {
	resp, err := cli.Post(jianshuAddBlogUrl, goz.Options{
		Headers: blog.headers(),
		JSON: map[string]interface{}{
			"notebook_id": noteId,
			"title":       article.Title,
			"at_bottom":   false,
		},
	})

	if err != nil {
		return
	}

	body, err := resp.GetBody()
	if err != nil {
		return
	}

	var respMap map[string]interface{}
	err = json.Unmarshal(body, &respMap)
	if err != nil {
		return
	}

	articleId = strconv.FormatFloat(respMap["id"].(float64), 'f', -1, 64)
	return
}

// 编辑
func (blog *JianShu) editArticle(articleId string, article util.Article) (err error) {
	resp, err := cli.Put(fmt.Sprintf(jianshuEditBlogUrl, articleId), goz.Options{
		Headers: blog.headers(),
		JSON: map[string]interface{}{
			"id":               articleId,
			"title":            article.Title,
			"content":          article.Content,
			"autosave_control": 2,
		},
	})

	if err != nil {
		return
	}

	body, err := resp.GetBody()
	if err != nil {
		return
	}

	var respMap map[string]interface{}
	err = json.Unmarshal(body, &respMap)
	if err != nil {
		return
	}

	return
}

// 发布
func (blog *JianShu) releaseArticle(articleId string) (err error) {
	_, err = cli.Post(fmt.Sprintf(jianshuReleaseBlogUrl, articleId), goz.Options{
		Headers: blog.headers(),
	})

	return
}

func (blog *JianShu) headers() map[string]interface{} {
	return map[string]interface{}{
		"User-Agent":   "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36",
		"Accept":       "application/json",
		"Content-Type": "application/json; charset=UTF-8",
		"Cookie":       viper.GetString("platform.jianshu.cookie"),
	}
}
