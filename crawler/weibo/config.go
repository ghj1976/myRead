package weibo

import (
	"github.com/BurntSushi/toml"
	"log"
	"os"
	"osext"
	"path"
)

// 微博开放平台的配置参数
type WeiboOpenAPIConfig struct {
	AppKey    string `toml:"key"`
	AppSecret string `toml:"secret"`
	Code      string `toml:"code"`
}

// 给指定位置文件写TOML文件
func WriteTOMLFile(fileName string, obj interface{}) {

	fo, err := os.Create(fileName)
	if err != nil {
		log.Println(err)
		return
	}
	defer fo.Close()

	// var firstBuffer bytes.Buffer
	e := toml.NewEncoder(fo)
	err = e.Encode(obj)
	if err != nil {
		log.Println(err)
		return
	}
}

// 初始化配置文件
func InitConfigFile() {

	var config WeiboOpenAPIConfig

	// config
	// http://open.weibo.com/apps  这里可以看到的的 AppKey  AppSecret 这两个值。
	// Code 应用授权码
	// 参考 http://open.weibo.com/wiki/%E6%8E%88%E6%9D%83%E6%9C%BA%E5%88%B6%E8%AF%B4%E6%98%8E#.E4.BD.BF.E7.94.A8OAuth2.0.E8.B0.83.E7.94.A8API
	config = WeiboOpenAPIConfig{AppKey: "AppKey",
		AppSecret: "AppSecret",
		Code:      "96dc8ae47a6960674f87af99a50687d4"}

	dir, err := osext.ExecutableFolder()
	if err != nil {
		log.Println("osext.ExecutableFolder()", err)
	}

	filename := path.Join(dir, "config.toml")
	log.Println(filename)

	WriteTOMLFile(filename, config)
}
