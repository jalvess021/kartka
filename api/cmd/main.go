package main

import (
	"encoding/json"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/jalvess021/kartka/api/internal/db" // Atualize o caminho conforme a sua estrutura de diretórios
	"github.com/jalvess021/kartka/api/internal/infra/akafka"
	repository "github.com/jalvess021/kartka/api/internal/infra/repository/pgsql"
	"github.com/jalvess021/kartka/api/internal/infra/web"
	usecase "github.com/jalvess021/kartka/api/internal/usecase/product"
)

func main() {

	db, err := db.ConnectDb();
	if  err != nil {
		panic(err)
	}
	defer db.Close();

	r := chi.NewRouter()

	//configurando cors para acesso do front-end
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://127.0.0.1:5173", "http://127.0.0.1:8181"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true, // Permite cookies e credenciais
	}))
	
	// chama as rotas
	web.SetupRoutes(r, db);
	go http.ListenAndServe(":8282", r)
	
	msgChan := make(chan *kafka.Message)
	go akafka.Consume([]string{"product"}, "kafka:9092", msgChan)

	//Cria repositorio e casos de usuo
	repository := repository.NewProductRepositoryPgsql(db)
	createProductUseCase := usecase.NewCreateProductUseCase(repository)

	//Acessando por dto
	for msg := range msgChan {
		dto := usecase.CreateProductInputDto{}
		err := json.Unmarshal(msg.Value, &dto)
		if err != nil {
			//logar erro
		}
		_, err = createProductUseCase.Execute(dto)
	}
}
