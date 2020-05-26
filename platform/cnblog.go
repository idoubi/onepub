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
	cnBlogUserInfoUrl string = "https://i.cnblogs.com/api/user"
)

// 博客园 https://www.cnblogs.com/
type CnBlog struct {
	info platInfo
}

// Login 模拟登陆
func (blog *CnBlog) Login() error {

	return nil
}

func (blog *CnBlog) IsLogin() error {
	resp, err := cli.Get(cnBlogUserInfoUrl, goz.Options{
		Headers: blog.headers(),
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

	if _, ok := respMap["loginName"]; !ok {
		return errors.New("cnblog login status verification failed")
	}

	return nil
}

// 发布
func (blog *CnBlog) Publish(article util.Article) error {
	uri := blog.info.publishURL

	resp, err := cli.Post(uri, goz.Options{
		Headers: blog.headers(),
		JSON: map[string]interface{}{
			"title":       article.Title,
			"postBody":    article.HtmlContent,
			"postType":    viper.GetInt("platform.cnblog.postType"),
			"isPublished": true,
		},
	})

	if err != nil {
		return err
	}

	_, err = resp.GetBody()
	if err != nil {
		return fmt.Errorf("cnblog publish error %w", err)
	}

	return nil
}

func (blog *CnBlog) headers() map[string]interface{} {
	return map[string]interface{}{
		"User-Agent":   "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36",
		"Content-Type": "application/json",
		"Cookie":       viper.GetString("platform.cnblog.cookie"),
		"x-blog-id":    viper.GetString("platform.cnblog.blog-id"),
	}
}
