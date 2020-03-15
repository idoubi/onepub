package platform

import (
	"encoding/json"
	"fmt"
	"github.com/idoubi/goz"
	"time"
)

type JianshuPlatInfo struct {
	getAllNotebooks       string
	createPostURLTemplate string
}
type Jianshu struct {
	info      platInfo
	extraInfo JianshuPlatInfo
}

func (jianshu *Jianshu) Login() error {
	return nil
}

// Login 模拟登陆
func (jianshu *Jianshu) Publish(title string, content string) error {
	cli := goz.NewClient()
	// list notebooks
	resp, _ := cli.Get("https://www.jianshu.com/author/notebooks", goz.Options{
		Headers: getJianshuHeaders(),
	})
	body, _ := resp.GetBody()
	var jsonMaps []map[string]interface{}
	_ = json.Unmarshal(body, &jsonMaps)
	// get default notebook id
	defaultNoteBookId := jsonMaps[0]["id"]

	// create a note in the default notebook
	data := map[string]interface{}{
		"notebook_id": defaultNoteBookId,
		// when creating, default title is today's date
		"title":     time.Now().Format("YYYY-MM-DD"),
		"at_bottom": true,
	}
	resp, _ = cli.Post("https://www.jianshu.com/author/notes", goz.Options{
		JSON:    data,
		Headers: getJianshuHeaders(),
	})
	body, _ = resp.GetBody()
	var jsonMap map[string]interface{}
	_ = json.Unmarshal(body, &jsonMap)
	createdNoteId := jsonMap["id"]

	// modify title and content of the new note
	data = map[string]interface{}{
		"id": createdNoteId,
		// to ensure the order
		"autosave_control": 1,
		"title":            title,
		"content":          content,
	}
	resp, _ = cli.PUT(fmt.Sprintf("https://www.jianshu.com/author/notes/%s", createdNoteId), goz.Options{
		JSON:    data,
		Headers: getJianshuHeaders(),
	})

	// publicise
	data = map[string]interface{}{}
	resp, _ = cli.Post("https://www.jianshu.com/author/notes/64727154/publicize", goz.Options{
		JSON:    data,
		Headers: getJianshuHeaders(),
	})
	return nil
}

// getCommonHeaders 生成通用请求头
func getJianshuHeaders() map[string]interface{} {
	return map[string]interface{}{
		"Cache-Control": "no-cache",
		"Accept":        "*/*",
		"User-Agent":    "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36",
		// TODO cookie filling
		"Cookie": "",
		"Host":   "www.jianshu.com",
	}
}
