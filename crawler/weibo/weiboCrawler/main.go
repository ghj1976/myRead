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
	} else {
		fmt.Println(jj)
	}

	fmt.Println(jj.AccessToken)
	tl, err := weibo.GetMyTimeline(jj.AccessToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(tl)
	//fmt.Println(len(tl.statuses))

}
