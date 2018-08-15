package main

import (
	"bytes"
	"encoding/json"
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

func sendRequest(url string, data string) (*http.Response, error) {
	fmt.Println(url)
	fmt.Println(data)
	d := []byte(data)
	req, err := http.NewRequest("POST", "http://127.0.0.1:5000", bytes.NewBuffer(d))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	return resp, err
}

type resStruct struct {
	Url  string
	Data json.RawMessage
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Body)
	decoder := json.NewDecoder(r.Body)
	var res resStruct
	err := decoder.Decode(&res)

	if err != nil {
		fmt.Println(err)
	}
	server, err := machinery.NewServer(cnf)
	if err != nil {
		fmt.Println(err)
	}

	strData, err := json.Marshal(res.Data)
	signature := &tasks.Signature{
		Name: "sendRequest",
		Args: []tasks.Arg{
			{
				Type:  "string",
				Value: res.Url,
			},
			{
				Type:  "string",
				Value: strData,
			},
		},
	}

	asyncResult, err := server.SendTask(signature)
	if err != nil {
		fmt.Println(err)
	}

	results, err := asyncResult.Get(time.Duration(time.Millisecond * 5))
	if err != nil {
		fmt.Errorf("Getting task result failed with error: %s", err.Error())
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
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func main() {
	server, err := machinery.NewServer(cnf)
	if err != nil {
		fmt.Println(err)
	}

	go runServer()
	server.RegisterTask("sendRequest", sendRequest)
	listen(server)
}
