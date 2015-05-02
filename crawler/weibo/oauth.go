package weibo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const (
	API_URL = "https://api.weibo.com/"
)

// 微博oauth验证必须的参数
type WeiboAuth struct {
	ClientId     string
	ClientSecret string
	RedirectUri  string
}

type AccessTokenResult struct {
	//参考
	// http://open.weibo.com/wiki/%E6%8E%88%E6%9D%83%E6%9C%BA%E5%88%B6%E8%AF%B4%E6%98%8E#.E4.BD.BF.E7.94.A8OAuth2.0.E8.B0.83.E7.94.A8API
	AccessToken string `json:"access_token"`
	RemindIn    string `json:"remind_in"`
	ExpiresIn   int    `json:"expires_in"`
	Uid         string `json:"uid"`
}

func (auth *WeiboAuth) getAuthorizeUrl() string {
	params := url.Values{}
	params.Add("client_id", auth.ClientId)
	params.Add("response_type", "code")
	params.Add("redirect_uri", auth.RedirectUri)
	return fmt.Sprintf("%soauth2/authorize?%s", API_URL, params.Encode())
}

func (auth *WeiboAuth) GetAccessToken(code string) (AccessTokenResult, error) {
	// https://api.weibo.com/oauth2/access_token?client_id=YOUR_CLIENT_ID&client_secret=YOUR_CLIENT_SECRET&grant_type=authorization_code&redirect_uri=YOUR_REGISTERED_REDIRECT_URI&code=CODE
	params := url.Values{}
	params.Add("client_id", auth.ClientId)
	params.Add("client_secret", auth.ClientSecret)
	params.Add("grant_type", "authorization_code")
	params.Add("redirect_uri", auth.RedirectUri)
	params.Add("code", code)
	accessTokenUrl := fmt.Sprintf("%soauth2/access_token?", API_URL)
	response, err := http.PostForm(accessTokenUrl, params)
	if err != nil {
		log.Println("error while get authorize code")
		panic(err)
	}
	defer response.Body.Close()
	accessTokenResult := AccessTokenResult{}
	str, _ := ioutil.ReadAll(response.Body)
	log.Printf("body:%s", str)
	err = json.Unmarshal(str, &accessTokenResult)
	if err != nil {
		log.Println("error while parsing token json")
		panic(err)
	}
	return accessTokenResult, nil
}
