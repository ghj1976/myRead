package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/ghj1976/myRead/crawler/weibo"
)

func main() {

	// config
	var config weibo.WeiboOpenAPIConfig

	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		fmt.Println(err)
		return
	}

	oauth := weibo.WeiboAuth{ClientId: config.AppKey, ClientSecret: config.AppSecret, RedirectUri: "http://www.teaduoduo.com/weibo"}

	jj, err := oauth.GetAccessToken(config.Code)
	if err != nil {
		fmt.Println(err)
		return
	}
	config.AccessToken = jj.AccessToken
	//	fmt.Println(jj.AccessToken)
	weibo.InitTomlShortFileName("config.toml", config)

	tl, err := weibo.GetMyTimeline(config.AccessToken)
	if err != nil {
		fmt.Println(err)

		return
	} else if tl.ErrorCode != 0 {
		// linux 下高亮输出技术参考： http://studygolang.com/articles/2500
		fmt.Printf("%c[1;40;32m", 0x1B)
		fmt.Println("weibo.GetMyTimeline error")
		fmt.Println(tl.ErrorCode)
		fmt.Println(tl.ErrorMsg)
		fmt.Printf(tl.RequestUrl)
		fmt.Printf("%c[0m\r\n", 0x1B)
		return
	}

	//  把抓去的数据存储 sqlite 数据库。

	fmt.Println(tl)
	fmt.Println(tl.ErrorCode)
	//fmt.Println(len(tl.statuses))

}
