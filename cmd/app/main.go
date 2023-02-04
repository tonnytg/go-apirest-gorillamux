package app

import (
	"database/sql"
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/go-chi/chi/v5"
	"github.com/tonnytg/go-apirest/internal/infra/akafka"
	"github.com/tonnytg/go-apirest/internal/infra/repository"
	"github.com/tonnytg/go-apirest/internal/infra/web"
	"github.com/tonnytg/go-apirest/internal/usecase"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(host.docker.internal:3306)/products")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repository := repository.NewProductRepositoryMysql(db)
	createProductUsecase := usecase.NewCreateProductUseCase(repository)
	listProductsUsecase := usecase.NewListProductsUseCase(repository)

	productHandlers := web.NewProductHandlers(createProductUsecase, listProductsUsecase)

	r := chi.NewRouter()
	r.Post("/products", productHandlers.CreateProductHandler)
	r.Get("/products", productHandlers.ListProductsHandler)

	go http.ListenAndServe(":8000", r)

	msgChan := make(chan *kafka.Message)
	go akafka.Consume([]string{"product"}, "host.docker.internal:9094", msgChan)

	for msg := range msgChan {
		dto := usecase.CreateProductInputDto{}
		err := json.Unmarshal(msg.Value, &dto)
		if err != nil {
			// logar o erro
		}
		_, err = createProductUsecase.Execute(dto)
	}
}
