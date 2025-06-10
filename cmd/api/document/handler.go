package document

import (
	"fmt"
	"net/http"
	"strconv"

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

func (h *Handler) FindDocumetByIdHandler(w http.ResponseWriter, r *http.Request) {

	id, err := getIdRequestParam(r)

	if err != nil {
		utils.WriteErrorResponse[any](w, http.StatusBadRequest, "Invalid Id format") // ToDo maybe return 404 for security purposes
		return
	}

	document := h.documentStore.FindById(id)
	if document == nil {
		utils.WriteErrorResponse[any](w, http.StatusNotFound, fmt.Sprintf("Unable to find document with id: %v", id))
		return
	}

	utils.WriteSuccessResponse(w, http.StatusOK, "SUCCESS", document)

}

func (h *Handler) FindAllDocumentsHandler(w http.ResponseWriter, r *http.Request) {
	utils.WriteSuccessResponse(w, http.StatusOK, "SUCCESS", h.documentStore.FindAll())
}

func (h *Handler) DeleteDocumentByIdHandler(w http.ResponseWriter, r *http.Request) {

	id, err := getIdRequestParam(r)
	if err != nil {
		utils.WriteErrorResponse[any](w, http.StatusBadRequest, "Invalid Id format") // ToDo maybe return 404 for security purposes
		return
	}

	if ok := h.documentStore.DeleteById(id); !ok {
		utils.WriteErrorResponse[any](w, http.StatusNotFound, fmt.Sprintf("Unable to find document with id: %v", id))
		return
	}

	utils.WriteSuccessResponse[any](w, http.StatusOK, "Document deleted", nil)
}

func getIdRequestParam(r *http.Request) (int, error) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return id, fmt.Errorf("Invalid Id format")
	}

	return id, nil
}
