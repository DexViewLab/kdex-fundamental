package netclient

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var client map[string]string

func TestRedis(t *testing.T) {
	client = map[string]string{"default": "127.0.0.1:6379"}
	InitRedisClients(client)

	cli := GetRedisClient()
	assert.NotNil(t, cli)

	cli1 := GetRedisClientByAddr("127.0.0.1:6379")
	assert.NotNil(t, cli1)
	cli2 := GetRedisClientByName("default")
	assert.NotNil(t, cli2)
	cli3 := GetAllRedisClients()
	assert.NotNil(t, cli3)
}
