package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"time"

	"github.com/IRSHIT033/E-comm-GO-/server/User_service/bootstrap"
	event "github.com/IRSHIT033/E-comm-GO-/server/listener_service/Event"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	bootstrap.Envinitializer()
	//try to connect to rabbitMQ
	rabbitConn, err := connect()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer rabbitConn.Close()

	//start listening for messages
	log.Println("listening for and consuming RabbitMQ messages...")

	//create consumer
	consumer, err := event.NewConsumer(rabbitConn)
	if err != nil {
		panic(err)
	}

	//watch the queue and consume events
	err = consumer.Listen([]string{"log.INFO", "log.WARNING", "log.ERROR"})

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	defer rabbitConn.Close()
}

func connect() (*amqp.Connection, error) {
	var counts int64
	var backOff = 1 * time.Second
	var connection *amqp.Connection
	amqp_server_uri := os.Getenv("AMQP_SERVER_URI")
	//don't continue until rabbitMQ is ready
	for {
		c, err := amqp.Dial(amqp_server_uri)
		if err != nil {
			fmt.Println("RabbitMQ not yet ready...")
			counts++
		} else {
			connection = c
			break
		}

		if counts > 5 {
			fmt.Println(err)
			return nil, err
		}

		backOff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		log.Println("backing off...")
		time.Sleep(backOff)
		continue
	}
	return connection, nil
}
