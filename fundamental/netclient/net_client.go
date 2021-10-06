package netclient

import (
	"log"
	"net"
	"net/http"
	"net/url"
	"sync"
	"time"

	"database/sql"

	_ "github.com/go-sql-driver/mysql" // mysql sucks...
)

var httpOnce sync.Once

var rpcOnce sync.Once
var mysqlOnce sync.Once

var netClient *http.Client
var mysqlClient *sql.DB

// ProxyIPPort ...
var ProxyIPPort string

// MysqlDSN ...
var MysqlDSN string

var rpcIPPort = "gamepoint.sogou:80"
var rpcUser, rpcPass = "joe", "333"

// GetHTTPClient returns the instance of a http.Client
func GetHTTPClient() *http.Client {
	httpOnce.Do(func() {

		var proxy = func(_ *http.Request) (*url.URL, error) {
			if ProxyIPPort == "" {
				return nil, nil
			}
			return url.Parse("http://" + ProxyIPPort)
		}

		var netTransport = &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   5 * time.Second,
				KeepAlive: 30 * time.Second,
			}).Dial,
			TLSHandshakeTimeout:   5 * time.Second,
			ResponseHeaderTimeout: 5 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			Proxy: proxy,
		}

		netClient = &http.Client{
			Timeout:   time.Second * 5,
			Transport: netTransport,
		}

	})

	return netClient
}

// GetMysqlClient 返回一个初始化好的mysql 客户端
func GetMysqlClient() *sql.DB {
	mysqlOnce.Do(func() {
		db, err := sql.Open("mysql", MysqlDSN)
		if nil != err {
			log.Println("get mysql client error:", err)
		}
		mysqlClient = db
	})

	return mysqlClient
}
