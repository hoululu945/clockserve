package amqp

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"serve/global"
	"strconv"
	"time"
)

var (
	channel              *amqp.Channel
	exchangeName         = "clock-exchange12"
	dlx_exchangeName     = "dlx_clock_exchange12"
	routeKey             = "pro-expir-clock-key12"
	dlx_queue            = "dlx_queue_clock_queue12"
	nomor_queue          = "clockP_queue12"
	deadLetterRoutingKey = "my_routing_dead_key12"
)

func ExchangeDeclare(connection *amqp.Connection) error {
	channel, _ = connection.Channel()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//正常交换机
	if err := channel.ExchangeDeclare(
		exchangeName,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	); err != nil {

		return fmt.Errorf("Exchange Declare: %s", err)

	}

	//死信交换机
	if err := channel.ExchangeDeclare(
		dlx_exchangeName,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	); err != nil {

		return fmt.Errorf("dlx Exchange Declare: %s", err)

	}
	return nil
}
func Publish(ClockId uint, expir time.Duration) error {
	pub_data := struct {
		ID  int
		exp int
	}{
		int(ClockId),
		int(expir),
	}
	marshal, err2 := json.Marshal(pub_data)
	if err2 != nil {
		return fmt.Errorf("转换 json %s", err2)
	}
	fmt.Println("---------------", expir.Milliseconds(), expir)
	ExchangeDeclare(global.RBBITMQ_CON)
	//nomor queue 声明
	args := amqp.Table{
		"x-dead-letter-exchange":    dlx_exchangeName,     // 设置死信交换机
		"x-dead-letter-routing-key": deadLetterRoutingKey, // 设置死信路由键
		//"x-message-ttl":             5000,
	}
	declare, err := channel.QueueDeclare(
		nomor_queue,
		true,
		false,
		false,
		false,
		args,
	)
	if err != nil {
		return fmt.Errorf("QueueDeclare:%s", err)

	}

	//死信交换机队列
	dlx_queueDeclare, err := channel.QueueDeclare(dlx_queue, true, false, false, false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("dlx queue declar%s", err)
	}
	err = channel.QueueBind(dlx_queueDeclare.Name, deadLetterRoutingKey, dlx_exchangeName, false, nil)
	if err != nil {
		return fmt.Errorf("dlx bind error %s", err)
	}

	//

	if err := channel.QueueBind(
		declare.Name,
		routeKey,
		exchangeName,
		false,
		nil,
	); err != nil {

		return fmt.Errorf("queue-bind:%s", err)

	}
	fmt.Println("mq-desc ", string(marshal), marshal)
	var a interface{}
	a = 1
	i, ok := a.(int)
	fmt.Println(i, ok)
	s := strconv.Itoa(int(expir.Milliseconds()))
	fmt.Println("*************", s)
	if err := channel.Publish(
		exchangeName,
		routeKey,
		false,
		false,
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     "text/plain",
			ContentEncoding: "",
			Body:            marshal,
			DeliveryMode:    amqp.Transient, // 1=non-persistent, 2=persistent
			Priority:        0,              // 0-9
			Expiration:      s,
		},
	); err != nil {
		fmt.Println("ppp", err)
		return fmt.Errorf("Publish:%s", err)

	}
	fmt.Println("发布成功amqp")
	return nil

}
