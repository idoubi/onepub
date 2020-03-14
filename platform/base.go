package platform

// Platform 发布文章的平台
type Platform interface {
	// 模拟登陆
	Login() error
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
	if plat == "juejin" {
		p := &Juejin{
			info: platInfo{
				host:       "https://juejin.im",
				loginURL:   "https://juejin.im/auth/type/email",
				publishURL: "https://post-storage-api-ms.juejin.im/v1/draftStorage",
			},
		}

		return p
	}

	return nil
}
