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
	csdnUserInfoUrl string = "https://download-console-api.csdn.net/v1/user/getUserBaseInfo"
)

// csdn https://www.csdn.net/
type Csdn struct {
	info platInfo
}

// Login 模拟登陆
func (c *Csdn) Login() error {
	return nil
}

func (c *Csdn) IsLogin() error {
	resp, err := cli.Get(csdnUserInfoUrl, goz.Options{
		Headers: c.headers(),
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

	code, ok := respMap["code"].(float64)
	if !ok {
		return errors.New("csdn login status verification failed")
	}

	if code != 200 {
		return fmt.Errorf("csdn login status verification failed %f", code)
	}

	return nil
}

// 发布
func (c *Csdn) Publish(article util.Article) error {
	uri := c.info.publishURL

	resp, err := cli.Post(uri, goz.Options{
		Headers: c.headers(),
		JSON: map[string]interface{}{
			"title":             article.Title,
			"Description":       article.Title,
			"authorized_status": false,
			"categories":        viper.GetString("platform.csdn.categories"),
			"content":           article.HtmlContent,
			"markdowncontent":   article.Content,
			"not_auto_saved":    "1",
			"original_link":     "",
			"resource_url":      "",
			"readType":          "public",
			"source":            "pc_mdeditor",
			"status":            0,
			"tags":              viper.GetString("platform.csdn.tags"),
			"type":              "original",
		},
	})

	if err != nil {
		return err
	}

	body, err := resp.GetBody()
	if err != nil {
		return fmt.Errorf("cnblog publish error %w", err)
	}

	fmt.Println(body)

	return nil
}

func (c *Csdn) headers() map[string]interface{} {
	return map[string]interface{}{
		"User-Agent":   "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36",
		"Content-Type": "application/json",
		"Cookie":       viper.GetString("platform.csdn.cookie"),
	}
}
