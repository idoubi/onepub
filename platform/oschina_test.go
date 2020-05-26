package platform

import (
	"github.com/idoubi/onepub/util"
	"testing"
)

func TestOsChina_IsLogin(t *testing.T) {
	initConfig("../config/pub.yaml")
	article := util.Article{
		Title:       "文章标题",
		Content:     "真的、这个编辑器是我见过最菜的。",
		HtmlContent: "<p>test 22333<p>",
	}
	t.Log(New("oschina").Publish(article))
}
