package platform

import (
	"fmt"
	"github.com/idoubi/goz"
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

func (blog *CnBlog) Publish(article Article) error {
	cli := goz.NewClient()
	uri := blog.info.publishURL

	resp, err := cli.Post(uri, goz.Options{
		Headers: blog.Headers(),
		JSON: map[string]interface{}{
			"title":    article.Title,
			"postBody": article.Content,
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

func (blog *CnBlog) Headers() map[string]interface{} {
	// todo 从配置文件中获取 cookie
	return map[string]interface{}{
		"User-Agent":   "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36",
		"Content-Type": "application/json",
		"Cookie":       "_ga=GA1.2.888615519.1584153138; __gads=ID=276d6ef367ab49eb:T=1584153138:S=ALNI_MYEp__dVZB5H8p-ydQEcHhkreJVhA; _gid=GA1.2.1920668042.1588320079; .Cnblogs.AspNetCore.Cookies=CfDJ8B9DwO68dQFBg9xIizKsC6SwKvOCGJasNYSpdMh6k0i8s1ZVZcPXPt7UyNHCEjP2lt3T6NtqtXcbW01LP1Bo90EgUzxiu5hUC88JHWOC-j7rLqm5DpE5dEvmn1e3IGesvq5uvrrZ4dm5VyBN_243uCfwGjNHhKSydbBiBwab2B5r-PVlWTwYobGtWPK6zAggp8s2zsepHYULYk0JSW7QoIEtlrVmwMUtUwdjqbyC_kvX-P6WcJyWpAfGBpcSNevmvx0y-kvQa5jZE1znulM6lxaIgSZan-XpCkdN_2xYjn5aid4w8NHeemmwJOjW0ZkqXxoDuXUJ_7K05WOHIw0T8pV8mgc_VDtmaZLOwxCpMJYM3U74mJ9xkwkDmAEPtSnKkK5Y_9qBKVII-8hsQw8dD1FDwWZBfYZRPDxUFGv25Er3Q-sI3fuDJz-m7RSpFt1UmQjNIN0lMgTCUPG84p9jSj8J1kmYEbg2PcHMkmvH2TEQcfgTDj8lAFoxO7U68-C2QftSYVHBMqA5ULipG3z36gstRCmM_mDB8rcMSL9V_HFX; .CNBlogsCookie=D58EA9DA353DC6C53B14DF19C25165F9BB54158FBB86D51296C251E1EE7709A842085FFCF5BDF03F7A57DBD48DF173CBB8864DDEFA6FA4F5200DE3ED99026DED4DCF7BB2B1579A1B2C000D93BF816BAD9B40CD6C; _gat=1",
		"x-blog-id":    "601044",
	}
}
