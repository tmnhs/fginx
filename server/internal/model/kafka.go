package model

//topic
const (
	KafkaMessageTopic = "hust-mall-message"

	HEATBEAT = "heatbeat"
	PONG     = "pong"

	// 消息类型，单聊或者群聊
	MESSAGETYPEUSER  = 1
	MESSAGETYPEGROUP = 2

	//消息内容类型
	MessageText  = 1
	MessageFile  = 2
	MessageOrder = 3
	//消息队列类型
	GOCHANNEL = "gochannel"
	KAFKA     = "kafka"
)
