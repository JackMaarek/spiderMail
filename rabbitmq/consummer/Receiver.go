package consummer

import (
	"fmt"
	"github.com/streadway/amqp"
	"os"
)

func receiveToRabbit() {
	url := os.Getenv("AMQP_URL")

	if url == "" {
		url = "amqp://user:bitmani@localhost:5672"
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
		fmt.Println("message received: " + string(msg.Body))
		msg.Ack(false)
	}

	defer connection.Close()
}
