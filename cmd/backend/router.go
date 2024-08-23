package main

import (
	"github.com/go-chi/chi/v5"
)

// Router functions
func defineRoutes(router chi.Router) {
	// API routes
	apiRouter := chi.NewRouter()

	// Health check and endpoints routes
	router.Get("/api", healthCheckHandler)
	router.Get("/api/v1", endpointsHandler)

	// Volume routes
	apiRouter.Route("/volumes", func(r chi.Router) {
		configureVolumeRoutes(r)
	})

	router.Mount("/api/v1", apiRouter)
}

// Configure volume routes
func configureVolumeRoutes(router chi.Router) {
	router.Get("/", listVolumesHandler)
	router.Get("/{volume_id:[0-9]+}", getVolumeHandler)

	// Chapter routes within volumes
	router.Route("/{volume_id:[0-9]+}/chapters", func(r chi.Router) {
		configureChapterRoutes(r)
	})
}

// Configure chapter routes
func configureChapterRoutes(router chi.Router) {
	router.Get("/", listChaptersHandler)
	router.Get("/{chapter_id:[0-9]+}", getChapterHandler)

	// Page routes within chapters
	router.Route("/{chapter_id:[0-9]+}/pages", func(r chi.Router) {
		configurePageRoutes(r)
	})
}

// Configure page routes
func configurePageRoutes(router chi.Router) {
	router.Get("/", listPagesHandler)
	router.Get("/{page_id:[0-9]+}", getPageHandler)

	// Panel routes within pages
	router.Route("/{page_id:[0-9]+}/panels", func(r chi.Router) {
		configurePanelRoutes(r)
	})
}

// Configure panel routes
func configurePanelRoutes(router chi.Router) {
	router.Get("/", listPanelsHandler)
	router.Get("/{panel_id:[0-9]+}", getPanelHandler)
	router.Get("/{panel_id:[0-9]+}/{width:[0-9]+}px", getPanelImageHandler)
}
