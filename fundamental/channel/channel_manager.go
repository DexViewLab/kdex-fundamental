package channel

import (
	"github.com/cskr/pubsub"
	"time"
)

const (
	capacity int = 10000
)

var manager Manager

type ManagerImpl struct {
	pubSub *pubsub.PubSub
	keyManagers []KeyManagerProxy
}

type Manager interface {
	ShutDown()
	Close(string)
	AddSubscribe(string, chan interface{})
	UnSubscribe(string, chan interface{})
	TryPublish(string, interface{})
	TrySendOnce(ch chan interface{}, data interface{})
	AddProxy(proxy ...KeyManagerProxy)
}

func InitChannelManager(cap int) Manager {
	c := capacity
	if cap != 0 {
		c = cap
	}
	manager = &ManagerImpl{
		pubSub: pubsub.New(c),
	}
	return manager
}

func GetChanelManager() Manager {
	return manager
}

func (c *ManagerImpl) ShutDown() {
	c.pubSub.Shutdown()
}

func (c *ManagerImpl) AddProxy(proxy ...KeyManagerProxy) {
	c.keyManagers = proxy

	go c.startPingTimer()
	go c.startMessageTimer()
}

func (c *ManagerImpl) Close(key string) {
	c.pubSub.Close(key)
}

func (c *ManagerImpl) AddSubscribe(key string, ch chan interface{}) {
	c.pubSub.AddSub(ch, key)
}

func (c *ManagerImpl) UnSubscribe(key string, ch chan interface{}) {
	c.pubSub.Unsub(ch, key)
}

func (c *ManagerImpl) TryPublish(key string, data interface{}) {
	c.pubSub.TryPub(data, key)
}

func (c *ManagerImpl) TrySendOnce(ch chan interface{}, data interface{}) {
	select {
	case ch <- data:
	default:
	}
}

func (c *ManagerImpl) startPingTimer() {
	ticker := time.NewTicker(30 * time.Second)
	for {
		for _, keyManager := range c.keyManagers {
			if keyManager != nil {
				keyManager.SendPingMessage()
			}
		}

		<-ticker.C
	}
}

func (c *ManagerImpl) startMessageTimer() {
	ticker := time.NewTicker(1 * time.Minute)
	for {
		for _, keyManager := range c.keyManagers {
			if keyManager != nil {
				keyManager.SendMessage()
			}
		}

		<-ticker.C
	}
}
