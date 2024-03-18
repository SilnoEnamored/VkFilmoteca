package main

import (
	"filmoteca/internal/db"
	"filmoteca/internal/handler"
	"filmoteca/internal/service"
	"log"
	"net/http"
)

func main() {
	// Инициализация подключения к базе данных
	database := db.New()
	defer database.Close()

	// Инициализация сервисов
	actorService := service.NewActorService(database)
	movieService := service.NewMovieService(database)

	// Инициализация обработчиков
	actorHandler := handler.NewActorHandler(actorService)
	movieHandler := handler.NewMovieHandler(movieService)

	// Настройка маршрутов
	http.HandleFunc("/actors", actorHandler.HandleActors)
	http.HandleFunc("/actors/", actorHandler.HandleActors) // Для обработки запросов с ID
	http.HandleFunc("/movies", movieHandler.HandleMovies)
	http.HandleFunc("/movies/", movieHandler.HandleMovies) // Для обработки запросов с ID

	// Запуск HTTP сервера
	log.Println("Server is running on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}
