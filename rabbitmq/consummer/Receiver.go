package consummer

import (
	"fmt"
	"github.com/streadway/amqp"
	"os"
)

func ReceiveToRabbit() {
	var messages []byte
	var messageList []byte

	url := os.Getenv("AMQP_URL")

	if url == "" {
		url = "amqp://user:guest@rabbitmq:5672"
	}
	// Connect to the rabbitMQ instance
	connection, err := amqp.Dial(url)

	if err != nil {
		panic("could not establish connection with RabbitMQ:" + err.Error())
	}

	channel, err := connection.Channel()

	if err != nil {
		panic("could not open RabbitMQ channel:" + err.Error())
	}

	msgs, err := channel.Consume("test", "", false, false, false, false, nil)

	if err != nil {
		panic("error consuming the queue: " + err.Error())
	}

	// The msgs will be a go channel, not an amqp channel
	for msg := range msgs {
		messageList = append(messages, msg.Body...)
		fmt.Println("message received: " + string(msg.Body))
		msg.Ack(false)
	}

	fmt.Println(messageList)
	for message := range messageList {
		fmt.Println(string(message))
	}

	//defer connection.Close()
}
