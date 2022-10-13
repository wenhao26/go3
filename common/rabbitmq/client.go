package rabbitmq

import (
	"fmt"

	"github.com/streadway/amqp"
)

type Option struct {
	Username string
	Password string
	Host     string
	Port     int64
}

type MQHandler struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

// 构建连接
func NewConn(opt Option) *MQHandler {
	url := fmt.Sprintf("amqp://%s:%s@%s:%d/", opt.Username, opt.Password, opt.Host, opt.Port)
	conn, err := amqp.Dial(url)
	if err != nil {
		panic(fmt.Errorf("MQ构建连接失败: %s \n", err))
	}

	ch, err := conn.Channel()
	if err != nil {
		panic(fmt.Errorf("MQ频道创建失败: %s \n", err))
	}
	// defer ch.Close()

	return &MQHandler{
		Connection: conn,
		Channel:    ch,
	}
}
