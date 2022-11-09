package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/danilodcn/go-intensivo/internal/order/infra/database"
	"github.com/danilodcn/go-intensivo/internal/order/usecase"
	"github.com/danilodcn/go-intensivo/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"

	// sqlite3
	_ "github.com/mattn/go-sqlite3"
)

func getConnection() *sql.DB {
	db, err := sql.Open("sqlite3", "./orders.db")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS orders (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL, PRIMARY KEY (id))")

	if err != nil {
		panic(err)
	}
	return db
}

func main() {

	db := getConnection()

	ch, err := rabbitmq.OpenChanel()
	if err != nil {
		panic(err)
	}

	repo := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPriceUseCase(repo)

	defer ch.Close()

	out := make(chan amqp.Delivery)

	go rabbitmq.Consume(ch, out)
	println("Consumer is running...\n")
	number := 0
	for msg := range out {
		var orderInputDTO usecase.OrderInputDTO
		// println(string(msg.Body), "\n")
		err := json.Unmarshal(msg.Body, &orderInputDTO)

		if err != nil {
			print("erro no json: ", err.Error(), "\n")
			continue
		}
		_, err = uc.Execute(&orderInputDTO)
		if err != nil {
			print("erro ao executar: ", err.Error(), "\n")
			continue
		}
		number++
		if number%10 == 0 {
			fmt.Println(orderInputDTO)
		}
		fmt.Println(number, " process messages")
		msg.Ack(false)
	}
}
