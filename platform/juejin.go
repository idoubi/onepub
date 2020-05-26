package platform

import (
	"encoding/json"
	"fmt"
	"github.com/idoubi/goz"
	"github.com/idoubi/onepub/util"
	"github.com/spf13/viper"
	"strconv"
)

func init() {
	cli = goz.NewClient()
}

var (
	cli            *goz.Request
	userInfoUrl    string = "https://juejin.im/auth"
	postPublishUrl string = "https://post-storage-api-ms.juejin.im/v1/postPublish"
	updateDraftUrl string = "https://post-storage-api-ms.juejin.im/v1/updateDraft"
)

// Juejin 掘金 https://Juejin.im
type Juejin struct {
	info platInfo
}

type jOpt struct {
	Token    string `json:"token"`
	ClientId int64  `json:"clientId"`
	UserId   string `json:"userId"`
}

// Login 模拟登陆
func (j *Juejin) Login() error {
	cli := goz.NewClient()
	uri := j.info.loginURL
	headers := getCommonHeaders()

	data := map[string]interface{}{
		"email":    "",
		"password": "",
	}

	resp, err := cli.Post(uri, goz.Options{
		Headers: headers,
		JSON:    data,
	})
	if err != nil {
		return err
	}

	body, err := resp.GetBody()
	if err != nil {
		return err
	}
	fmt.Println(body)
	return nil
}

func (j *Juejin) IsLogin() error {
	_, err := j.getOpt()
	if err != nil {
		return fmt.Errorf("juejin login status err: %w", err)
	}

	return nil
}

func (j *Juejin) Publish(article util.Article) error {
	opt, err := j.getOpt()
	if err != nil {
		return err
	}

	articleId, err := j.draftStorage(article, opt)
	if err != nil {
		return err
	}

	return j.postArticle(articleId, opt)
}

func (j *Juejin) headers() map[string]interface{} {
	// todo 从配置文件中获取 cookie
	return map[string]interface{}{
		"User-Agent":   "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36",
		"Content-Type": "application/x-www-form-urlencoded",
		"Cookie":       viper.GetString("platform.juejin.cookie"),
	}
}

func (j *Juejin) getOpt() (opt jOpt, err error) {
	resp, err := cli.Get(userInfoUrl, goz.Options{
		Headers: j.headers(),
	})
	if err != nil {
		return
	}

	body, err := resp.GetBody()
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &opt)
	if err != nil {
		err = fmt.Errorf("juejin: get auth error:%s", body)
		return
	}

	return
}

// 添加随笔
func (j *Juejin) draftStorage(article util.Article, opt jOpt) (articleId string, err error) {
	resp, err := cli.Post(j.info.publishURL, goz.Options{
		Headers: j.headers(),
		FormParams: map[string]interface{}{
			"uid":                    opt.UserId,
			"token":                  opt.Token,
			"device_id":              strconv.FormatInt(opt.ClientId, 10),
			"src":                    "web",
			"type":                   "markdown",
			"title":                  article.Title,
			"screenshot":             "",
			"markdown":               article.Content,
			"isTitleImageFullscreen": "0",
			"html":                   article.HtmlContent,
			"content":                "",
			"category":               viper.GetString("platform.juejin.category"),
			"tags":                   viper.GetString("platform.juejin.tags"),
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

	data, err := util.Sliceconv(respMap["d"])
	if err != nil {
		return
	}

	articleId = data[0].(string)
	return
}

// 更新随笔
func (j *Juejin) updateStorage(articleId string, article util.Article, opt jOpt) (err error) {
	resp, err := cli.Post(updateDraftUrl, goz.Options{
		Headers: j.headers(),
		FormParams: map[string]interface{}{
			"uid":                    opt.UserId,
			"token":                  opt.Token,
			"device_id":              strconv.FormatInt(opt.ClientId, 10),
			"src":                    "web",
			"type":                   "markdown",
			"title":                  article.Title,
			"screenshot":             "",
			"markdown":               article.Content,
			"isTitleImageFullscreen": "0",
			"html":                   "",
			"content":                "",
			"category":               viper.GetString("platform.juejin.category"),
			"tags":                   viper.GetString("platform.juejin.tags"),
		},
	})

	if err != nil {
		return
	}

	_, err = resp.GetBody()
	return
}

// 发布随笔
func (j *Juejin) postArticle(articleId string, opt jOpt) error {
	_, err := cli.Post(postPublishUrl, goz.Options{
		Headers: j.headers(),
		FormParams: map[string]interface{}{
			"uid":       opt.UserId,
			"token":     opt.Token,
			"device_id": strconv.FormatInt(opt.ClientId, 10),
			"src":       "web",
			"postId":    articleId,
		},
	})

	return err
}

// getCommonHeaders 生成通用请求头
func getCommonHeaders() map[string]interface{} {
	return map[string]interface{}{
		"User-Agent":   "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36",
		"Content-Type": "application/json",
	}
}
