package main

import (
	"fmt"
	"github.com/ghj1976/myRead/crawler/weibo"
)

func main() {
	oauth := weibo.WeiboAuth{ClientId: "***", ClientSecret: "***", RedirectUri: "http://www.teaduoduo.com/weibo"}

	jj, err := oauth.GetAccessToken("****")
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
	} else {
		fmt.Println(tl)
	}

}
