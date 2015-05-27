package weibo

// sqlite 数据库相关操作

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"strings"
)

var (

	// 数据库连接字符串
	DBFilePath = "./db/proxy.db"
	//DBFilePath = "/root/proxy/db/proxy.db"
)

// 获得所有代理服务器的地址
func GetProxyList() (arr []string, isOK bool) {
	arr = make([]string, 0)
	// 连接数据库
	db, err := sql.Open("sqlite3", DBFilePath)
	defer db.Close()
	if err != nil {
		return nil, false
	}

	rows, err := db.Query(`SELECT proxyUrl FROM proxy order by responseTime asc,lastupdate desc;`)
	if err != nil {
		return nil, false
	}

	for rows.Next() {
		var url string
		err = rows.Scan(&url)
		if err == nil {
			arr = append(arr, url)
		}
	}
	return arr, true
}

// 只有在不存在的时候才插入, 如果存在，则返回false
func InsertProxyInfo(url string) bool {
	// 连接数据库
	db, err := sql.Open("sqlite3", DBFilePath)
	defer db.Close()
	if err != nil {
		log.Println(err)
		return false
	}

	stmt, err := db.Prepare(`insert into proxy (proxyUrl,isHideIP,ipAdd,lastupdate,responseTime) 
values(?,?,?,datetime('now', 'localtime'),?) `)
	if err != nil {
		log.Println(err)
		return false
	}

	_, err = stmt.Exec(url, false, "", 5000)
	if err != nil {
		if !strings.Contains(err.Error(), "proxyurl_UNIQUE") {
			log.Println(err)
		}
		return false
	}

	log.Println("inserted：", url)

	return true
}

// 超时的数据更新逻辑
func UpdateTimeOutProxyInfo(url string, responseTime int64) bool {
	// 连接数据库
	db, err := sql.Open("sqlite3", DBFilePath)
	defer db.Close()
	if err != nil {
		log.Println(err)
		return false
	}

	stmt, err := db.Prepare(`update proxy set responseTime = ?,lastupdate = datetime('now', 'localtime') where proxyUrl = ? `)
	if err != nil {
		log.Println(err)
		return false
	}

	_, err = stmt.Exec(responseTime, url)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
