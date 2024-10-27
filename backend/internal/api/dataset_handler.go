package api

import (
	"context"
	"data-merging/internal/service"
	"encoding/json"
	"net/http"
	"time"
)

// DatasetHandler handles HTTP requests for profiles
type DatasetHandler struct {
	datasetService service.DatasetService
}

func NewDatasetHandler(datasetService service.DatasetService) *DatasetHandler {
	return &DatasetHandler{datasetService: datasetService}
}

// GetDatasetHandler handles GET requests to fetch profiles
func (h *DatasetHandler) GetDatasetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	profiles, err := h.datasetService.GetDatasets(ctx)
	if err != nil {
		http.Error(w, "Unable to fetch datasets", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(profiles)
}
