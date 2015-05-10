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

	// oauth := weibo.WeiboAuth{ClientId: config.AppKey, ClientSecret: config.AppSecret, RedirectUri: "http://www.teaduoduo.com/weibo"}

	// jj, err := oauth.GetAccessToken(config.Code)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// config.AccessToken = jj.AccessToken
	// fmt.Println(jj.AccessToken)

	tl, err := weibo.GetMyTimeline(config.AccessToken)
	if err != nil {
		fmt.Println(err)

		return
	} else if tl.ErrorCode != 0 {
		fmt.Println("weibo.GetMyTimeline error")
		fmt.Println(tl.ErrorCode)
		fmt.Println(tl.ErrorMsg)
		fmt.Println(tl.RequestUrl)
		return
	}

	fmt.Println(tl)
	fmt.Println(tl.ErrorCode)
	//fmt.Println(len(tl.statuses))

}
