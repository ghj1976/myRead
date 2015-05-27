package weibo

import (
	"encoding/json"
	"fmt"
	"time"
)

// 微博内容相关的实体类
type tweet struct {
	CreateAt           time.Time        `json:"create_at"`               // 发表时间 微博创建时间
	Id                 int64            `json:"id"`                      // 编号 微博ID
	Mid                string           `json:"mid"`                     //
	Idstr              string           `json:"idstr"`                   //
	Text               string           `json:"text"`                    // 内容 微博信息内容
	SourceAllowclick   int              `json:"source_allowclick"`       //
	SourceType         int              `json:"source_type"`             //
	Source             string           `json:"source"`                  //
	Favorited          bool             `json:"favorited"`               //
	Truncated          bool             `json:"truncated"`               //
	InReply2StatusId   string           `json:"in_reply_to_status_id"`   //
	InReply2UserId     string           `json:"in_reply_to_user_id"`     //
	InReply2ScreenName string           `json:"in_reply_to_screen_name"` //
	PicUrls            []string         `json:"pic_urls"`                //
	Geo                string           `json:"geo"`                     //
	Uid                int64            `json:"uid"`                     //
	Pid                int64            `json:"pid"`                     //
	ThumbnailPic       string           `json:"thumbnail_pic"`           // 缩略图片地址，没有时不返回此字段
	BmiddlePic         string           `json:"bmiddle_pic"`             // 中等尺寸图片地址，没有时不返回此字段
	OriginalPic        string           `json:"original_pic"`            // 原始图片地址，没有时不返回此字段
	PicIds             []string         `json:"pic_ids"`                 // 微博配图ID。多图时返回多图ID，用来拼接图片url。用返回字段thumbnail_pic的地址配上该返回字段的图片ID，即可得到多个图片url
	RepostsCount       int              `json:"reposts_count"`           //
	CommentsCount      int              `json:"comments_count"`          //
	AttitudesCount     int              `json:"attitudes_count"`         //
	MLevel             int              `json:"mlevel"`                  //
	Visible            visibleStruct    `json:"visible"`                 //
	DarwinTag          darwinTagsStruct `json:"darwin_tag"`              //
}

type visibleStruct struct {
	VType  int `json:"type"`
	ListId int `json:"list_id"`
}

type darwinTagsStruct struct {
}

type WeiboTopic struct {
	tweet
	RetweetedStatus tweet `json:"retweeted_status"` // 被转发的微博
}

type WeiboTimeline struct {
	Statuses       []WeiboTopic `json:"statises"`        // 微博内容集合
	TotalNumber    int          `json:"total_number"`    // 总数
	HasVisible     bool         `json:"hasvisible"`      //
	Interval       int          `json:"interval"`        //
	NextCursor     int          `json:"next_cursor"`     //
	PreviousCursor int          `json:"previous_cursor"` //
	ErrorMsg       string       `json:"error"`           //
	ErrorCode      int          `json:"error_code"`      //
	RequestUrl     string       `json:"request"`         //
}

// 参考接口文档：
// http://open.weibo.com/wiki/Statuses/user_timeline
// 这里使用下面工具做json 格式化： http://json.parser.online.fr/
func GetMyTimeline(access_token string) (timeline WeiboTimeline, err error) {
	tl := new(WeiboTimeline)

	url := "https://api.weibo.com/2/statuses/user_timeline.json"
	params := map[string]interface{}{"trim_user": 1, "access_token": access_token}
	bbody, err := HttpGet(url, access_token, params)
	if err != nil {
		fmt.Println(err)
		return *tl, err
	}
	WriteFile2ExecutableFolder2("1.txt", bbody)
	// fmt.Println(string(bbody))

	// fmt.Println("")

	err = json.Unmarshal(bbody, tl)
	if err != nil {
		fmt.Println("...")
		fmt.Println(err)
		return *tl, err
	}
	return *tl, nil
}
