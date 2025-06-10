package document

import (
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {

	mux.HandleFunc("POST /documents", h.CreateDocumentHandler)
	mux.HandleFunc("GET /documents/{id}", h.FindDocumetByIdHandler)
	mux.HandleFunc("GET /documents", h.FindAllDocumentsHandler)
	mux.HandleFunc("DELETE /documents/{id}", h.DeleteDocumentByIdHandler)
	mux.HandleFunc("GET /documents/search", h.FilterDocumentsHandler)

}
