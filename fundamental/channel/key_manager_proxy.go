package channel

type KeyManagerProxy interface {
	SendPingMessage()
	SendMessage()
}
