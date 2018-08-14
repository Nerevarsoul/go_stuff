package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/RichardKnop/machinery/v1/tasks"
)

var cnf = &config.Config{
	Broker:        "amqp://rabbitmq:rabbitmq@172.26.0.3:5672/",
	DefaultQueue:  "machinery_tasks",
	ResultBackend: "amqp://rabbitmq:rabbitmq@172.26.0.3:5672/",
	AMQP: &config.AMQPConfig{
		Exchange:     "machinery_exchange",
		ExchangeType: "direct",
		BindingKey:   "machinery_task",
	},
}

func Add(args ...int64) (int64, error) {
	sum := int64(0)
	for _, arg := range args {
		sum += arg
	}
	return sum, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	server, err := machinery.NewServer(cnf)
	if err != nil {
		fmt.Println(err)
	}

	signature := &tasks.Signature{
		Name: "add",
		Args: []tasks.Arg{
			{
				Type:  "int64",
				Value: 1,
			},
			{
				Type:  "int64",
				Value: 1,
			},
		},
	}

	asyncResult, err_r := server.SendTask(signature)
	if err_r != nil {
		fmt.Println(err_r)
	}

	results, err_as := asyncResult.Get(time.Duration(time.Millisecond * 5))
	if err_as != nil {
		fmt.Errorf("Getting task result failed with error: %s", err_as.Error())
	}
	fmt.Println(tasks.HumanReadableResults(results))
}

func listen(server *machinery.Server) {
	worker := server.NewWorker("worker_name", 10)
	err := worker.Launch()
	if err != nil {
		fmt.Println(err)
	}
}

func runServer() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":9002", nil))
}

func main() {
	server, err := machinery.NewServer(cnf)
	if err != nil {
		fmt.Println(err)
	}

	go runServer()
	server.RegisterTask("add", Add)
	listen(server)
}
