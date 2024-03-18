package handler

import (
	"encoding/json"
	"filmoteca/internal/models"
	"filmoteca/internal/service"
	"net/http"
	"strconv"
	"strings"
)

type MovieHandler struct {
	Service *service.MovieService
}

func NewMovieHandler(service *service.MovieService) *MovieHandler {
	return &MovieHandler{
		Service: service,
	}
}

func (h *MovieHandler) HandleMovies(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getMovies(w, r)
	case http.MethodPost:
		h.createMovie(w, r)
	case http.MethodPut:
		h.updateMovie(w, r)
	case http.MethodDelete:
		h.deleteMovie(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *MovieHandler) getMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := h.Service.GetAllMovies(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func (h *MovieHandler) createMovie(w http.ResponseWriter, r *http.Request) {
	var movie models.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.Service.CreateMovie(r.Context(), &movie); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(movie)
}

func (h *MovieHandler) updateMovie(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/movies/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid movie ID", http.StatusBadRequest)
		return
	}

	var movie models.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.Service.UpdateMovie(r.Context(), id, &movie); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movie)
}

func (h *MovieHandler) deleteMovie(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/movies/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid movie ID", http.StatusBadRequest)
		return
	}

	if err := h.Service.DeleteMovie(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
