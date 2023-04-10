package routes

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	
	"github.com/bootcamp-go/desafio-go-web/internal/tickets"
)

type Response struct {
	Success bool     `json:"success"`
	Data interface{} `json:"data"`
	Error string     `json:"error"`
}

func handleGetByCountry(repo tickets.Repository) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		
		dest := r.URL.Query().Get("dest")
		
		if dest == "" {
			http.Error(w, "Missing destination parameter", http.StatusBadRequest)
			return
		}
		
		ticketList, err := repo.GetTicketByDestination(ctx, dest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		response := Response {
			Success: true,
			Data: ticketList,
			Error: "",
		}
		
		json.NewEncoder(w).Encode(response)
	}
}

func handleGetAverage(repo tickets.Repository) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		
		dest := r.URL.Query().Get("dest")
		
		if dest == "" {
			http.Error(w, "Missing destination parameter", http.StatusBadRequest)
			return
		}
		
		service := tickets.NewService(repo)
		
		avg, err := service.AverageDestination(ctx, dest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		response := Response {
			Success: true,
			Data: avg,
			Error: "",
		}
		json.NewEncoder(w).Encode(response)
	}
}
		
