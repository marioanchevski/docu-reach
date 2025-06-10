package document

import (
	"net/http"

	"github.com/marioanchevski/docu-reach/types"
	"github.com/marioanchevski/docu-reach/utils"
)

type Handler struct {
	documentStore types.DocumentStore
}

func NewHandler(ds types.DocumentStore) *Handler {
	return &Handler{
		documentStore: ds,
	}
}

func (h *Handler) CreateDocumentHandler(w http.ResponseWriter, r *http.Request) {
	var payload types.CreateDocumentRequest

	if err := utils.ReadJSON(w, r, &payload); err != nil {
		utils.WriteErrorResponse[any](w, http.StatusBadRequest, err.Error())
		return
	}

	if err := payload.Validate(); err != nil {
		utils.WriteErrorResponse[any](w, http.StatusBadRequest, err.Error())
		return
	}

	response := h.documentStore.Create(payload)
	utils.WriteSuccessResponse(w, http.StatusCreated, "SUCCESS", response)
}
