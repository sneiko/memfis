package handlers

import (
	"encoding/json"
	"net/http"

	"memfis/internal/models"
	"memfis/internal/parser"
	"memfis/templates"
)

type Handler struct {
	memoryData *models.MemoryData
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) LoadFile(filename string) error {
	data, err := parser.ParseMemoryData(filename)
	if err != nil {
		return err
	}
	h.memoryData = data
	return nil
}

func (h *Handler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	if h.memoryData == nil {
		// Show file upload interface
		component := templates.FileUpload()
		component.Render(r.Context(), w)
		return
	}

	// Show dashboard
	component := templates.Dashboard(h.memoryData)
	component.Render(r.Context(), w)
}

func (h *Handler) APIDataHandler(w http.ResponseWriter, r *http.Request) {
	if h.memoryData == nil {
		http.Error(w, "No data loaded", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(h.memoryData)
}

func (h *Handler) APIMemStatsHandler(w http.ResponseWriter, r *http.Request) {
	if h.memoryData == nil {
		http.Error(w, "No data loaded", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(h.memoryData.MemStats)
}

func (h *Handler) APIStackTracesHandler(w http.ResponseWriter, r *http.Request) {
	if h.memoryData == nil {
		http.Error(w, "No data loaded", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(h.memoryData.StackTraces)
}
