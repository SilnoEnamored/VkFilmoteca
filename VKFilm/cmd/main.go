package main

import (
	"filmoteca/internal/db"
	"filmoteca/internal/handler"
	"filmoteca/internal/service"
	"log"
	"net/http"
)

func main() {
	database := db.New()
	defer database.Close()

	//сервисы
	actorService := service.NewActorService(database)
	movieService := service.NewMovieService(database)

	//ручки
	actorHandler := handler.NewActorHandler(actorService)
	movieHandler := handler.NewMovieHandler(movieService)

	//маршруты
	http.HandleFunc("/actors", actorHandler.HandleActors)
	http.HandleFunc("/actors/", actorHandler.HandleActors) // Для обработки запросов с ID
	http.HandleFunc("/movies", movieHandler.HandleMovies)
	http.HandleFunc("/movies/", movieHandler.HandleMovies) // Для обработки запросов с ID

	//запуск сервера
	log.Println("Server is running on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}
