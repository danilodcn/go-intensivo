package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/danilodcn/go-intensivo/internal/order/infra/database"
	"github.com/danilodcn/go-intensivo/internal/order/usecase"
	"github.com/danilodcn/go-intensivo/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"

	// sqlite3
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/lib/pq"
)

func getConnection() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@db:5432/postgres?sslmode=disable")
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
	total := make(chan int)

	go rabbitmq.Consume(ch, out)
	time.Sleep(10 * time.Millisecond)
	println("Consumer is running...")
	numberOfWorkers := 80
	for i := 1; i <= numberOfWorkers; i++ {
		go worker(out, total, uc, i)
	}

	go getTotal(total)
	forever := make(chan int)

	<-forever
}

func getTotal(total chan int) {
	println("get total is running ...")
	n := 0
	for i := range total {
		n += i
		println(n)
	}
}

func worker(out chan amqp.Delivery, total chan int, uc *usecase.CalculateFinalPriceUseCase, id int) {
	workerName := fmt.Sprint("[WORKER ", id, "] ")
	println(workerName, " is running...")
	number := 0

	for msg := range out {
		var orderInputDTO usecase.OrderInputDTO
		err := json.Unmarshal(msg.Body, &orderInputDTO)

		if err != nil {
			println("erro no json: ", err.Error())
			continue
		}
		_, err = uc.Execute(&orderInputDTO)
		if err != nil {
			println("erro ao executar: ", err.Error())
			fmt.Println(orderInputDTO)
			// panic(err)
			msg.Ack(false)
			continue
		}
		number++
		if number%10 == 0 {
			fmt.Println(workerName, orderInputDTO)
		}
		msg.Ack(false)

		total <- 1
		time.Sleep(time.Millisecond * 0)
	}
}
