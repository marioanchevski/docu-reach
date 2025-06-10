package health

import (
	"net/http"

	"github.com/marioanchevski/docu-reach/utils"
)

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {

		utils.WriteHealthResponse(w)

	})
}
