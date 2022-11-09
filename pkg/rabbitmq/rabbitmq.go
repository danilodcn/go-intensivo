package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func OpenChanel() (*amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672")
	if err != nil {
		print("Erro na conexão: ", string(err.Error()))
		panic(err)
	}

	ch, err := conn.Channel()

	if err != nil {
		print("Erro ao abrir o canal: ", string(err.Error()))
		panic(err)
	}
	return ch, err
}

func Consume(ch *amqp.Channel, out chan amqp.Delivery) error {
	msgs, err := ch.Consume("orders", "go-consume", false, false, false, false, nil)
	print("[Consume] is starting ...\n")
	if err != nil {
		println("Erro de conexão: ", err.Error())
		panic(err)
	}
	for msg := range msgs {
		// print("new message found: ", string(msg.Body), "\n")
		out <- msg
	}
	return nil
}
