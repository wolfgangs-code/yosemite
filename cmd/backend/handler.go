package main

// ? See router rules in router.go

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// ? Handler functions
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"status": "OK"}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func endpointsHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"status": "OK"}
	w.WriteHeader(http.StatusServiceUnavailable)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// ! Volumes
func listVolumesHandler(w http.ResponseWriter, r *http.Request) {
	response := []map[string]string{
		{"volume_id": "1", "title": "Dummy Volume 1"},
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func getVolumeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	volumeID := vars["volume_id"]
	response := map[string]string{"volume_id": volumeID, "title": "Volume Title"}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// ! Chapters
func listChaptersHandler(w http.ResponseWriter, r *http.Request) {
	response := []map[string]string{
		{"chapter_id": "1", "title": "Dummy Chapter 1"},
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func getChapterHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	chapterID := vars["chapter_id"]
	response := map[string]string{"chapter_id": chapterID, "title": "Chapter Title"}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// ! Pages
func listPagesHandler(w http.ResponseWriter, r *http.Request) {
	response := []map[string]string{
		{"page_id": "1", "title": "Dummy Page 1"},
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func getPageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageID := vars["page_id"]
	response := map[string]string{"page_id": pageID, "title": "Page Title"}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// ! Panels
func listPanelsHandler(w http.ResponseWriter, r *http.Request) {
	response := []map[string]string{
		{"panel_id": "1", "title": "Dummy Panel 1"},
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func getPanelHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	panelID := vars["panel_id"]
	response := map[string]string{"panel_id": panelID, "title": "Panel Title"}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func getPanelImageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	panelID := vars["panel_id"]
	width := vars["width"]
	response := map[string]string{"panel_id": panelID, "title": "Panel Title", "width": width}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
