package platform

import (
	"fmt"

	"github.com/idoubi/goz"
)

// Juejin 掘金 https://Juejin.im
type Juejin struct {
	info platInfo
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

func (j *Juejin) Publish(article Article) error {
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

// getCommonHeaders 生成通用请求头
func getCommonHeaders() map[string]interface{} {
	return map[string]interface{}{
		"User-Agent":   "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36",
		"Content-Type": "application/json",
	}
}
