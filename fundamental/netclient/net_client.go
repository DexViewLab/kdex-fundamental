package netclient

import (
	"log"
	"net"
	"net/http"
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


// GetHTTPClient returns the instance of a http.Client
func GetHTTPClient() *http.Client {
	httpOnce.Do(func() {
		var netTransport = &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   10 * time.Second,
				KeepAlive: 30 * time.Second,
			}).Dial,
			TLSHandshakeTimeout:   10 * time.Second,
			ResponseHeaderTimeout: 10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		}

		netClient = &http.Client{
			Timeout:   time.Second * 10,
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
