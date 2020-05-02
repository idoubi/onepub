package platform

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
}

// Platform 发布文章的平台
type Platform interface {
	// 模拟登陆
	Login() error
	Publish(article Article) error
}

// platInfo 平台信息
type platInfo struct {
	host       string // 域名
	loginURL   string // 登陆地址
	username   string // 登陆账号
	password   string // 登陆密码
	publishURL string // 发布地址
}

// 博客信息
type Article struct {
	Title   string
	Content string
}

// New 初始化
func New(plat string) Platform {
	if platform, ok := platformInfo[plat]; ok {
		return platform
	}

	return nil
}
