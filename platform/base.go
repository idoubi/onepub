package platform

import (
	"github.com/idoubi/onepub/util"
)

var platformInfo map[string]Platform = map[string]Platform{
	"juejin": &Juejin{
		info: platInfo{
			host:       "https://juejin.im",
			loginURL:   "https://juejin.im/auth/type/email",
			publishURL: "https://post-storage-api-ms.juejin.im/v1/draftStorage",
		},
	},
	"cnblog": &CnBlog{
		info: platInfo{
			host:       "",
			loginURL:   "",
			publishURL: "https://i-beta.cnblogs.com/api/posts",
		},
	},
	"jianshu": &JianShu{
		info: platInfo{
			host:       "",
			loginURL:   "",
			publishURL: "",
		},
	},
	"oschina": &OsChina{
		info: platInfo{
			host:       "",
			loginURL:   "",
			publishURL: "",
		},
	},
}

// Platform 发布文章的平台
type Platform interface {
	// 模拟登陆
	Login() error
	// 检验登录态
	IsLogin() error
	// 发布
	Publish(article util.Article) error
}

// platInfo 平台信息
type platInfo struct {
	host       string // 域名
	loginURL   string // 登陆地址
	username   string // 登陆账号
	password   string // 登陆密码
	publishURL string // 发布地址
}

// New 初始化
func New(plat string) Platform {
	if platform, ok := platformInfo[plat]; ok {
		return platform
	}

	return nil
}

// 所有的平台名
func AllPlatform() []string {
	keys := make([]string, 0, len(platformInfo))
	for k := range platformInfo {
		keys = append(keys, k)
	}
	return keys
}
