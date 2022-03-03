package handler

import (
	"log"
	"context"
	"net/http"
	"encoding/json"
	"github.com/TechBowl-japan/go-stations/model"
	"github.com/TechBowl-japan/go-stations/service"
)

// A TODOHandler implements handling REST endpoints.
type TODOHandler struct {
	svc *service.TODOService
}

// NewTODOHandler returns TODOHandler based http.Handler.
func NewTODOHandler(svc *service.TODOService) *TODOHandler {
	return &TODOHandler{
		svc: svc,
	}
}

// Create handles the endpoint that creates the TODO.
func (h *TODOHandler) Create(ctx context.Context, req *model.CreateTODORequest) (*model.CreateTODOResponse, error) {
	create, err := h.svc.CreateTODO(ctx, req.Subject, req.Description)
	if err != nil{
		log.Println(err)
	}
	response := &model.CreateTODOResponse{*create}
	return response, nil
}

// Read handles the endpoint that reads the TODOs.
func (h *TODOHandler) Read(ctx context.Context, req *model.ReadTODORequest) (*model.ReadTODOResponse, error) {
	_, _ = h.svc.ReadTODO(ctx, 0, 0)
	return &model.ReadTODOResponse{}, nil
}

// Update handles the endpoint that updates the TODO.
func (h *TODOHandler) Update(ctx context.Context, req *model.UpdateTODORequest) (*model.UpdateTODOResponse, error) {
	_, _ = h.svc.UpdateTODO(ctx, 0, "", "")
	return &model.UpdateTODOResponse{}, nil
}

// Delete handles the endpoint that deletes the TODOs.
func (h *TODOHandler) Delete(ctx context.Context, req *model.DeleteTODORequest) (*model.DeleteTODOResponse, error) {
	_ = h.svc.DeleteTODO(ctx, nil)
	return &model.DeleteTODOResponse{}, nil
}


func (h *TODOHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	switch r.Method {
	case "POST":
		var req *model.CreateTODORequest
		json.NewDecoder(r.Body).Decode(&req)
		if req.Subject == "" {
			w.WriteHeader(http.StatusBadRequest)
		}else{
			create, err := h.Create(r.Context(), req)
			response, err := json.Marshal(create)
			if err != nil {
				log.Println(err)
			}
			w.WriteHeader(http.StatusOK)
			w.Write(response)
		}
	}
}