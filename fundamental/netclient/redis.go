package netclient

import (
	"log"
	"net/url"
	"regexp"
	"strconv"
	"sync"

	"github.com/go-redis/redis"
	"github.com/golang/glog"
)

var redisOnce sync.Once
var redisClients map[string]*redis.Client

// RedisIPPort ...
var RedisIPPort = "127.0.0.1:6379"

const (
	defaultRedisName = "default"
)

func safeRedisAddr(addr string) string {
	return regexp.MustCompile(`://([^:]+):(.*)@`).ReplaceAllString(addr, "://$1:****@")
}

// InitRedisClients 初始化一次，一次性把所有客户端弄好
func InitRedisClients(addrs map[string]string) {
	redisOnce.Do(func() {
		glog.V(8).Info("netclient InitRedisClients.")
		redisClients = make(map[string]*redis.Client)
		// create every client by addr
		for name, addr := range addrs {
			createOneRedisClient(name, addr)
		}
	})
}

// GetRedisClient returns the instance of a redis.Client
func GetRedisClient() *redis.Client {
	redisOnce.Do(func() {
		// 如果没有调用过InitRedisClients，说明是老代码，那么我们初始化一个默认的client放到map里去
		redisClients = make(map[string]*redis.Client)
		createOneRedisClient(defaultRedisName, RedisIPPort)
	})
	return GetRedisClientByName(defaultRedisName)
}

// GetRedisClientByName 根据name来找client，如果name非法直接panic掉
func GetRedisClientByName(name string) *redis.Client {
	if _, exists := redisClients[name]; !exists {
		log.Fatalf("try to get redis which doesn't not exists. name = %s", name)
	}
	return redisClients[name]
}

// GetAllRedisClients 返回所有初始化过的clientss
func GetAllRedisClients() []*redis.Client {
	var clients []*redis.Client
	for _, c := range redisClients {
		clients = append(clients, c)
	}
	return clients
}

// redis://db:password@host:port
func parseRedisAddr(addr string) (host string, password string, db int) {
	db = 0
	u, err := url.Parse(addr)
	if err != nil {
		host = addr
	} else {
		host = u.Host
		db64, _ := strconv.ParseInt(u.User.Username(), 0, 32)

		db = int(db64)
		password, _ = u.User.Password()
	}

	glog.V(8).Infof("parse redis URI. addr = %s, host = %s, db = %d", safeRedisAddr(addr), host, db)

	return
}

func createOneRedisClient(name, addr string) {
	host, password, db := parseRedisAddr(addr)
	redisClients[name] = redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password, // no password set
		DB:       db,       // use default DB
	})
}

// IsRedisNil 判断Redis是否是空
func IsRedisNil(err error) bool {
	return err == redis.Nil
}

// InitRedisClientsUseAlias 初始化一次，一次性把所有客户端弄好
//用于 ut-docker，docker中的redis 地址为例   "trade_center": "test:windtest@tcp(mysql:3306)/trade_center?charset=utf8",
func InitRedisClientsUseAlias(addrs map[string]string) {
	redisOnce.Do(func() {
		glog.V(8).Infof("netclient InitRedisClients.")
		redisClients = make(map[string]*redis.Client)
		// create every client by addr
		for name, addr := range addrs {
			createOneRedisClientUseAlias(name, addr)
		}
	})
}

//该 创建redis客户端方法区别于createOneRedisClient ，去掉了parseRedisAddr，
//用于 ut-docker
func createOneRedisClientUseAlias(name, addr string) {
	//host, password, db := parseRedisAddr(addr)
	redisClients[name] = redis.NewClient(&redis.Options{
		Addr: addr,
		//Password: password, // no password set
		//DB:       db,       // use default DB
	})
}

// GetRedisClientByAddr 单纯的根据addr得到client
func GetRedisClientByAddr(addr string) *redis.Client {
	host, password, db := parseRedisAddr(addr)
	return redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       db,
	})
}
