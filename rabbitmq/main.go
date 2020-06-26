package rabbitmq

import "github.com/JackMaarek/spiderMail/rabbitmq/consummer"

func main()  {
	consummer.ReceiveFromRabbit()
}
