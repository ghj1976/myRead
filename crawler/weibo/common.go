package weibo

import (
	"compress/gzip"
	"errors"
	"fmt"
	"github.com/gosexy/to"
	"io/ioutil"
	"net/http"
	"net/url"
)

// 读取 gzip 加密的内容
func read_body(response *http.Response) ([]byte, error) {

	switch response.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err := gzip.NewReader(response.Body)
		if err != nil {
			return nil, err
		}
		defer reader.Close()
		contents, err := ioutil.ReadAll(reader)
		if err != nil {
			return nil, err
		}
		return contents, nil
	default:
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		return contents, nil
	}

	return nil, errors.New("Unknow Errors")
}

// 组成http 要传递的参数
func encodeParams(params map[string]interface{}) (string, error) {
	if len(params) > 0 {
		values := url.Values{}
		for key, value := range params {
			values.Add(key, to.String(value))
		}
		return values.Encode(), nil
	}
	return "", errors.New("Params Is Empty!")
}

func HttpGet(the_url string, authorization string, params map[string]interface{}) ([]byte, error) {

	url_params, err := encodeParams(params)
	http_url := fmt.Sprintf("%v?%v", the_url, url_params)
	fmt.Println(http_url)

	request, err := http.NewRequest("GET", http_url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Accept-Encoding", "gzip")

	// if authorization != "" {
	// 	request.Header.Add("Authorization", fmt.Sprintf("OAuth2 %s", authorization))
	// }

	client := &http.Client{}
	response, err := client.Do(request) // Do Request
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := read_body(response)
	if err != nil {
		return nil, err
	}

	return body, nil
}
