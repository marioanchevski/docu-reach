package document

import (
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {

	mux.HandleFunc("POST /documents", h.CreateDocumentHandler)
}
