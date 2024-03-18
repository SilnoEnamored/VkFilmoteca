package handler

import (
	"encoding/json"
	"filmoteca/internal/models"
	"filmoteca/internal/service"
	"net/http"
	"strconv"
	"strings"
)

type ActorHandler struct {
	Service *service.ActorService
}

func NewActorHandler(service *service.ActorService) *ActorHandler {
	return &ActorHandler{
		Service: service,
	}
}

func (h *ActorHandler) HandleActors(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getActors(w, r)
	case http.MethodPost:
		h.createActor(w, r)
	case http.MethodPut:
		h.updateActor(w, r)
	case http.MethodDelete:
		h.deleteActor(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *ActorHandler) getActors(w http.ResponseWriter, r *http.Request) {
	actors, err := h.Service.GetAllActors(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(actors)
}

func (h *ActorHandler) createActor(w http.ResponseWriter, r *http.Request) {
	var actor models.Actor
	if err := json.NewDecoder(r.Body).Decode(&actor); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.Service.CreateActor(r.Context(), &actor); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(actor)
}

func (h *ActorHandler) updateActor(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/actors/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid actor ID", http.StatusBadRequest)
		return
	}

	var actor models.Actor
	if err := json.NewDecoder(r.Body).Decode(&actor); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.Service.UpdateActor(r.Context(), id, &actor); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(actor)
}

func (h *ActorHandler) deleteActor(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/actors/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid actor ID", http.StatusBadRequest)
		return
	}

	if err := h.Service.DeleteActor(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
