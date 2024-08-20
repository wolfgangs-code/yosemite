package main

// ? See handler rules in handler.go

import (
	"github.com/gorilla/mux"
)

// ? Router functions
func defineRoutes(router *mux.Router) {
	// API routes
	apiRouter := router.PathPrefix("/api/v1").Subrouter()

	// Health check and endpoints routes
	router.HandleFunc("/api", healthCheckHandler).Methods("GET")
	router.HandleFunc("/api/v1", endpointsHandler).Methods("GET")

	// Volume routes
	volumeRouter := apiRouter.PathPrefix("/volumes").Subrouter()
	configureVolumeRoutes(volumeRouter)
}

// ! Configure volume routes
func configureVolumeRoutes(router *mux.Router) {
	router.HandleFunc("", listVolumesHandler).Methods("GET")
	router.HandleFunc("/{volume_id:[0-9]+}", getVolumeHandler).Methods("GET")

	// Chapter routes within volumes
	chapterRouter := router.PathPrefix("/{volume_id:[0-9]+}/chapters").Subrouter()
	configureChapterRoutes(chapterRouter)
}

// ! Configure chapter routes
func configureChapterRoutes(router *mux.Router) {
	router.HandleFunc("", listChaptersHandler).Methods("GET")
	router.HandleFunc("/{chapter_id:[0-9]+}", getChapterHandler).Methods("GET")

	// Page routes within chapters
	pageRouter := router.PathPrefix("/{chapter_id:[0-9]+}/pages").Subrouter()
	configurePageRoutes(pageRouter)
}

// ! Configure page routes
func configurePageRoutes(router *mux.Router) {
	router.HandleFunc("", listPagesHandler).Methods("GET")
	router.HandleFunc("/{page_id:[0-9]+}", getPageHandler).Methods("GET")

	// Panel routes within pages
	panelRouter := router.PathPrefix("/{page_id:[0-9]+}/panels").Subrouter()
	configurePanelRoutes(panelRouter)
}

// ! Configure panel routes
func configurePanelRoutes(router *mux.Router) {
	router.HandleFunc("", listPanelsHandler).Methods("GET")
	router.HandleFunc("/{panel_id:[0-9]+}", getPanelHandler).Methods("GET")
	router.HandleFunc("/{panel_id:[0-9]+}/{width:[0-9]+}px", getPanelImageHandler).Methods("GET")
}
