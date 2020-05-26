package platform

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/idoubi/goz"
	"github.com/idoubi/onepub/util"
	"github.com/spf13/viper"
)

var (
	osChinaFindAllMsgCount string = "https://my.oschina.net/u/%s/msg/findAllMsgCount"
	osChinaReleaseBlogUrl  string = "https://my.oschina.net/u/%s/blog/save"
)

// 开源中国 https://www.oschina.net/
type OsChina struct {
	info platInfo
}

// Login 模拟登陆
func (blog *OsChina) Login() error {
	return nil
}

func (blog *OsChina) IsLogin() error {
	resp, err := cli.Post(fmt.Sprintf(osChinaFindAllMsgCount, viper.GetString("platform.oschina.user_id")), goz.Options{
		Headers: blog.headers(),
		FormParams: map[string]interface{}{
			"user_code": viper.GetString("platform.oschina.user_code"),
			"userId":    viper.GetString("platform.oschina.user_id"),
		},
	})

	if err != nil {
		return err
	}

	body, err := resp.GetBody()
	if err != nil {
		return err
	}

	var respMap map[string]interface{}
	err = json.Unmarshal(body, &respMap)
	if err != nil {
		return err
	}

	if int64(respMap["code"].(float64)) != 1 {
		return errors.New(respMap["message"].(string))
	}

	return nil
}

// 发布
func (blog *OsChina) Publish(article util.Article) error {
	_, err := blog.releaseArticle(article)
	if err != nil {
		return fmt.Errorf("oschina add article error: %w", err)
	}

	return nil
}

// release
func (blog *OsChina) releaseArticle(article util.Article) (articleId string, err error) {
	resp, err := cli.Post(fmt.Sprintf(osChinaReleaseBlogUrl, viper.GetString("platform.oschina.user_id")), goz.Options{
		Headers: blog.headers(),
		FormParams: map[string]interface{}{
			"as_top":         viper.GetString("platform.oschina.as_top"),
			"catalog":        viper.GetString("platform.oschina.catalog"),
			"classification": viper.GetString("platform.oschina.classification"),
			"deny_comment":   viper.GetString("platform.oschina.deny_comment"),
			"privacy":        viper.GetString("platform.oschina.privacy"),
			"user_code":      viper.GetString("platform.oschina.user_code"),
			"downloadImg":    0,
			"type":           1,
			"isRecommend":    0,
			"content_type":   3,
			"title":          article.Title,
			"content":        article.Content,
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

	if int64(respMap["code"].(float64)) != 1 {
		err = errors.New(respMap["message"].(string))
		return
	}

	return
}

func (blog *OsChina) headers() map[string]interface{} {
	return map[string]interface{}{
		"User-Agent":       "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36",
		"Accept":           "application/json, text/javascript, */*; q=0.01",
		"Content-Type":     "application/x-www-form-urlencoded; charset=UTF-8",
		"Cookie":           viper.GetString("platform.oschina.cookie"),
		"X-Requested-With": "XMLHttpRequest",
	}
}
