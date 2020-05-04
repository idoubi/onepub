package platform

import (
	"fmt"
	"github.com/idoubi/goz"
	"github.com/idoubi/onepub/util"
	"github.com/spf13/viper"
)

// 博客园 https://www.cnblogs.com/
type CnBlog struct {
	info platInfo
}

// Login 模拟登陆
func (j *CnBlog) Login() error {
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

// 发布
func (blog *CnBlog) Publish(article util.Article) error {
	cli := goz.NewClient()
	uri := blog.info.publishURL

	resp, err := cli.Post(uri, goz.Options{
		Headers: blog.headers(),
		JSON: map[string]interface{}{
			"title":    article.Title,
			"postBody": article.HtmlContent,
			"postType": 1,
		},
	})

	if err != nil {
		return err
	}

	_, err = resp.GetBody()
	if err != nil {
		return err
	}

	return nil
}

func (blog *CnBlog) headers() map[string]interface{} {
	// todo 从配置文件中获取 cookie
	return map[string]interface{}{
		"User-Agent":   "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36",
		"Content-Type": "application/json",
		"Cookie":       viper.GetString("platform.cnblog.cookie"),
		"x-blog-id":    viper.GetString("platform.cnblog.blog-id"),
	}
}
