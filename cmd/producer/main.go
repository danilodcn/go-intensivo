package main

import (
	"encoding/json"
	"math/rand"
	"time"

	"github.com/danilodcn/go-intensivo/internal/order/entity"
	"github.com/danilodcn/go-intensivo/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"

	// sqlite3
	_ "github.com/mattn/go-sqlite3"
)

func Producer(ch *amqp.Channel, order entity.Order) error {
	body, err := json.Marshal(order)
	if err != nil {
		return err
	}
	err = ch.Publish(
		"amq.direct",
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	return err
}

func GenerateOrder() entity.Order {
	order, err := entity.NewOrder(
		rand.Float64()*100,
		rand.Float64()*30,
	)
	if err != nil {
		panic(err)
	}
	return *order
}

func main() {
	ch, err := rabbitmq.OpenChanel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	numberOfMessages := 2000
	println("generating", numberOfMessages, "orders")
	for i := 1; i <= numberOfMessages; i++ {
		err := Producer(ch, GenerateOrder())
		if err != nil {
			println(err.Error())
		}
		println(i+1, "messages generated")

		time.Sleep(time.Millisecond * 1)
	}
}
