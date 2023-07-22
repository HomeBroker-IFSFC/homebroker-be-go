package main

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/JoaoRafa19/homebroker/go/internal/infra/kafka"
	"github.com/JoaoRafa19/homebroker/go/internal/market/dto"
	"github.com/JoaoRafa19/homebroker/go/internal/market/entity"
	"github.com/JoaoRafa19/homebroker/go/internal/market/transformer"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	ordersIn := make(chan *entity.Order)
	ordersOut := make(chan *entity.Order)
	wg := &sync.WaitGroup{}
	defer wg.Wait()

	kafkaMsgChannel := make(chan *ckafka.Message)

	configMap := &ckafka.ConfigMap{
		"bootstrap.servers":"host.docker.internal:9094", 
		"group.id": "myGroup",
		"auto.offset.reset":"earliest",
	}

	producer := kafka.NewKafkaProducer(configMap)
	kafka := kafka.NewConsumer(configMap, []string{"input"})

	go kafka.Consume(kafkaMsgChannel) // Thread 2

	// recebe do canal do Kafka e joga no input, processa joga no output e publica no kafka
	book := entity.NewBook(ordersIn, ordersOut, wg)
	go book.Trade() // Thread 3

	go func () {
		for msg := range kafkaMsgChannel {
			wg.Add(1)
			tradeInput := dto.TradeInput {}
			err := json.Unmarshal(msg.Value, &tradeInput)
			if err == nil {
				panic(err)
			}
			order:= transformer.TransformInput(tradeInput)
			ordersIn <- order
		}
	} ()


	for res := range ordersOut {
		output := transformer.TransformOutput(res)
		outputJson,err := json.MarshalIndent( output, "", "    ")
		if err != nil {
			fmt.Println(err)
		}
		producer.Publish(outputJson, []byte("orders"), "output")
	}
}