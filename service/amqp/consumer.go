package amqp

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"serve/global"
	"serve/model"
	"serve/service/common"
	"time"
)

func NewConsumer() error {
	fmt.Println("开始进入clock consumer")
	channel, _ := global.RBBITMQ_CON.Channel()
	ch := struct {
		conn  *amqp.Connection
		chann *amqp.Channel
		done  chan error
	}{
		global.RBBITMQ_CON,
		channel,
		nil,
	}
	err := ch.chann.ExchangeDeclare(dlx_exchangeName, "direct", true, false, false, false, nil)
	if err != nil {
		return fmt.Errorf("declare exchange error %s", err)
	}
	queueDeclare, err := ch.chann.QueueDeclare(dlx_queue, true, false, false, false, nil)
	if err != nil {
		return fmt.Errorf("queue declare%s", err)
	}
	err = ch.chann.QueueBind(queueDeclare.Name, deadLetterRoutingKey, dlx_exchangeName, false, nil)
	if err != nil {
		return fmt.Errorf("queue bind %s", err)
	}

	consume, err := ch.chann.Consume(queueDeclare.Name, "clock-simple-consumer", false, false, false, false, nil)
	if err != nil {
		return fmt.Errorf("consumer %s", err)
	}
	go func(delivery <-chan amqp.Delivery, done chan error) {
		fmt.Println("开始消费")
		for d := range delivery {
			fmt.Println("消费", string(d.Body))
			log.Printf(
				"got %dB delivery: [%v] %q",
				len(d.Body),
				d.DeliveryTag,
				d.Body,
			)
			var bodyMap map[string]interface{}
			err2 := json.Unmarshal(d.Body, &bodyMap)
			if err2 != nil {
				fmt.Printf("pull err %s", err2)
			}
			fmt.Println(bodyMap)
			clock := model.Clocks{}

			global.Backend_DB.First(&clock, map[string]interface{}{"id": bodyMap["ID"]})
			common.CommonService.PubClock(&clock)
			clock.IsTip = 1
			global.Backend_DB.Save(clock)
			common.SendMail(new(common.WyMail), &clock)
			CircleSet(&clock)
			d.Ack(false)
		}
		log.Printf("handle: deliveries channel closed")
		done <- nil
	}(consume, ch.done)
	log.Printf("running forever")
	return nil

}
func CircleSet(clock *model.Clocks) {
	reminderType := clock.ReminderType
	if reminderType != 0 && clock.IsCircle == 1 {
		duration := time.Second
		switch reminderType {
		case 1:
			duration = 24 * 60 * 60 * time.Second
		case 2:
			duration = 24 * 60 * 60 * 7 * time.Second
		case 3:
			duration = 24 * 60 * 60 * 30 * time.Second
		case 4:
			duration = 24 * 60 * 60 * 365 * time.Second
		}
		err1 := Publish(clock.ID, duration)
		if err1 == nil {
			fmt.Println("wy添加新的循环成功成功！", duration)
		}

	}
	clock.IsTip = 1
	global.Backend_DB.Save(clock)

}
