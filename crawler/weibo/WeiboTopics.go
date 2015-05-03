package weibo

import (
	"encoding/json"
	"fmt"
	"time"
)

// 微博内容相关的实体类
type tweet struct {
	create_at               time.Time // 发表时间 微博创建时间
	id                      int64     // 编号 微博ID
	mid                     string    //
	idstr                   string    //
	text                    string    // 内容 微博信息内容
	source_allowclick       int       //
	source_type             int       //
	source                  string    //
	favorited               bool      //
	truncated               bool      //
	in_reply_to_status_id   string    //
	in_reply_to_user_id     string    //
	in_reply_to_screen_name string    //
	pic_urls                []string  //
	geo                     string
	uid                     int64    //
	pid                     int64    //
	thumbnail_pic           string   // 缩略图片地址，没有时不返回此字段
	bmiddle_pic             string   // 中等尺寸图片地址，没有时不返回此字段
	original_pic            string   // 原始图片地址，没有时不返回此字段
	pic_ids                 []string // 微博配图ID。多图时返回多图ID，用来拼接图片url。用返回字段thumbnail_pic的地址配上该返回字段的图片ID，即可得到多个图片url
	reposts_count           int
	comments_count          int
	attitudes_count         int
	mlevel                  int
}

type WeiboTopic struct {
	tweet
	retweeted tweet // 被转发的微博
}

type WeiboTimeline struct {
	statuses        []WeiboTopic // 微博内容集合
	total_number    int          // 总数
	hasvisible      bool         //
	interval        int          //
	next_cursor     int          //
	previous_cursor int          //
}

// 参考接口文档：
// http://open.weibo.com/wiki/Statuses/user_timeline
// 这里使用下面工具做json 格式化： http://json.parser.online.fr/
func GetMyTimeline(access_token string) (timeline WeiboTimeline, err error) {
	url := "https://api.weibo.com/2/statuses/user_timeline.json"
	params := map[string]interface{}{"trim_user": 1, "access_token": access_token}
	bbody, err := HttpGet(url, access_token, params)
	if err != nil {
		fmt.Println(err)
		return timeline, err
	}
	WriteFile2ExecutableFolder2("1.txt", bbody)
	fmt.Println(string(bbody))

	err = json.Unmarshal(bbody, &timeline)
	if err != nil {
		fmt.Println(err)
		return timeline, err
	}
	return timeline, nil
}
