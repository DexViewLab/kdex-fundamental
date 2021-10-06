package channel

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

type Message struct {
	key     string
	index   int
	message string
}

func publish(key string) {
	for i := 0; i < 50; i++ {
		msg := &Message{
			key:     key,
			index:   i,
			message: fmt.Sprintf("hello-%s-%d", key, i),
		}
		if i == 25 {
			msg = nil
		}

		GetChanelManager().TryPublish(key, msg)
		//time.Sleep(time.Duration(rand.RandInt64(100, 500)) * time.Millisecond)
	}
}

func receiveMsg(t *testing.T, ch chan interface{}, key string) {
	var msgs []*Message
loop:
	for {
		select {
		case rawData, open := <-ch:
			if !open {
				//fmt.Println("chan closed")
				break loop
			}
			msg := rawData.(*Message)
			msgs = append(msgs, msg)
			if msg == nil {
				//fmt.Println("receive empty message")
				break
			}

			assert.Equal(t, key, msg.key)
		}
	}

	assert.Equal(t, int(50), len(msgs))
}

func Test_ChannelManager(t *testing.T) {
	//InitChannelManager()

	key1 := "test1"
	key2 := "test2"
	key3 := "test3"
	var chTest1 []chan interface{}
	var chTest2 []chan interface{}
	var chTest3 []chan interface{}

	for i := 0; i < 10; i++ {
		ch := make(chan interface{}, 10000)
		chTest1 = append(chTest1, ch)
		GetChanelManager().AddSubscribe(key1, ch)
	}

	for i := 0; i < 20; i++ {
		ch := make(chan interface{}, 10000)
		chTest2 = append(chTest2, ch)
		GetChanelManager().AddSubscribe(key2, ch)
	}

	for i := 0; i < 30; i++ {
		ch := make(chan interface{}, 10000)
		chTest3 = append(chTest3, ch)
		GetChanelManager().AddSubscribe(key3, ch)
	}

	var wait sync.WaitGroup
	for _, ch := range chTest1 {
		wait.Add(1)
		go func() {
			receiveMsg(t, ch, key1)
			wait.Done()
		}()
		time.Sleep(100 * time.Millisecond)
	}

	for _, ch := range chTest2 {
		wait.Add(1)
		go func() {
			receiveMsg(t, ch, key2)
			wait.Done()
		}()
		time.Sleep(100 * time.Millisecond)
	}

	for _, ch := range chTest3 {
		wait.Add(1)
		go func() {
			receiveMsg(t, ch, key3)
			wait.Done()
		}()
		time.Sleep(100 * time.Millisecond)
	}

	time.Sleep(2 * time.Second)

	var wait2 sync.WaitGroup
	wait2.Add(1)
	go func() {
		publish(key1)
		wait2.Done()
	}()

	wait2.Add(1)
	go func() {
		publish(key2)
		wait2.Done()
	}()

	wait2.Add(1)
	go func() {
		publish(key3)
		wait2.Done()
	}()

	wait2.Wait()
	GetChanelManager().ShutDown()

	wait.Wait()
}
