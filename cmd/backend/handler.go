package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Helper function to handle responses
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payload)
}

// Generalized handler for fetching an item by ID
func getItemByID(w http.ResponseWriter, r *http.Request, idKey string, itemName string) {
	vars := mux.Vars(r)
	id := vars[idKey]
	response := map[string]string{idKey: id, "title": itemName}
	respondJSON(w, http.StatusOK, response)
}

// Generalized handler for listing items
func listItems(w http.ResponseWriter, items []map[string]string) {
	respondJSON(w, http.StatusOK, items)
}

// ? Handler functions
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"status": "OK"}
	respondJSON(w, http.StatusOK, response)
}

func endpointsHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"status": "OK"}
	respondJSON(w, http.StatusServiceUnavailable, response)
}

// ! Volumes
func listVolumesHandler(w http.ResponseWriter, r *http.Request) {
	volumes := []map[string]string{
		{"volume_id": "1", "title": "Dummy Volume 1"},
	}
	listItems(w, volumes)
}

func getVolumeHandler(w http.ResponseWriter, r *http.Request) {
	getItemByID(w, r, "volume_id", "Volume Title")
}

// ! Chapters
func listChaptersHandler(w http.ResponseWriter, r *http.Request) {
	chapters := []map[string]string{
		{"chapter_id": "1", "title": "Dummy Chapter 1"},
	}
	listItems(w, chapters)
}

func getChapterHandler(w http.ResponseWriter, r *http.Request) {
	getItemByID(w, r, "chapter_id", "Chapter Title")
}

// ! Pages
func listPagesHandler(w http.ResponseWriter, r *http.Request) {
	pages := []map[string]string{
		{"page_id": "1", "title": "Dummy Page 1"},
	}
	listItems(w, pages)
}

func getPageHandler(w http.ResponseWriter, r *http.Request) {
	getItemByID(w, r, "page_id", "Page Title")
}

// ! Panels
func listPanelsHandler(w http.ResponseWriter, r *http.Request) {
	panels := []map[string]string{
		{"panel_id": "1", "title": "Dummy Panel 1"},
	}
	listItems(w, panels)
}

func getPanelHandler(w http.ResponseWriter, r *http.Request) {
	getItemByID(w, r, "panel_id", "Panel Title")
}

func getPanelImageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	panelID := vars["panel_id"]
	width := vars["width"]
	response := map[string]string{"panel_id": panelID, "title": "Panel Title", "width": width}
	respondJSON(w, http.StatusOK, response)
}
