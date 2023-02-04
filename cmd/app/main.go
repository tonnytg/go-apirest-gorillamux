package app

import (
	"database/sql"
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/tonnytg/go-apirest/internal/infra/akafka"
	"github.com/tonnytg/go-apirest/internal/usecase"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(host.docker.internal:3306/products")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	msgChan := make(chan *kafka.Message)
	go akafka.Consume([]string{"products"}, "host.docker.internal:9094", msgChan)

	repository := repository.NewCreateProductUJseCase(repository)

	for msg := range msgChan {
		dto := usecase.CreateProductInputDto{}
		err := json.Unmarshal(msg.Value, &dto)
		if err != nil {
			log.Default().Println(err)
		}

		_, err = createProductUsecase.Execute(dto)
	}
}
